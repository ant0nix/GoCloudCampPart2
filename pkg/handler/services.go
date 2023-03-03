package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"

	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/gin-gonic/gin"
)

var playlist gocloudcamppart2.Playlist = *gocloudcamppart2.NewPlaylist()

func (h *Handler) Start(c *gin.Context) {

	playlist.Played = true
	var id int
	for {
		playlist.LastTrack = h.services.StartPLaylist.LastSong()
		playlist.FirstId = h.services.StartPLaylist.FirstSong()
		if playlist.Paused {
			playlist.Cond.Broadcast()
			playlist.Paused = false
		}

		opChan := make(chan string)

		if id < 0 {
			id = 0
			playlist.TrackID = 0
		}

		track, err := h.services.StartPLaylist.NextSong(id)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			log.Println(err.Error())
		}

		for track.ID == -1 {
			id++
			track, err = h.services.StartPLaylist.NextSong(id)

			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				log.Println(err.Error())
			}

		}

		go func(opChan chan string, track gocloudcamppart2.Track) {
			playlist.TrackID = id
			playlist.Play(opChan, track)

		}(opChan, track)

		c.Stream(func(w io.Writer) bool {
			output, ok := <-opChan
			if !ok {
				return false
			}
			outputBytes := bytes.NewBufferString(output)
			c.Writer.Write(append(outputBytes.Bytes(), []byte("\n")...))
			return true
		})
		if id == playlist.LastTrack || playlist.Stoped {
			playlist.Stoped = false
			break
		}
		if playlist.PrevTrack {
			track.ID = -1
			for track.ID == -1 {
				id--
				track, err = h.services.StartPLaylist.NextSong(id)
				if err != nil {
					newErrorResponse(c, http.StatusInternalServerError, err.Error())
					log.Println(err.Error())
				}
				if id == playlist.FirstId || id == 1 {
					break
				}
			}
			playlist.PrevTrack = false
		} else {
			id++
		}

	}

	playlist.Played = false
	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": "END",
	})

}

func (h *Handler) Next(c *gin.Context) {
	if playlist.Played {
		playlist.NextTrack = true
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "next track",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "playlist isn't playing",
		})
	}
}

func (h *Handler) Prev(c *gin.Context) {
	if playlist.Played {
		playlist.PrevTrack = true
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "prev track",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "playlist isn't playing",
		})
	}
}

func (h *Handler) Pause(c *gin.Context) {

	if playlist.Played {
		if !playlist.Paused {
			playlist.Paused = true
			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "paused track",
			})
		} else {

			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "playlist already paused",
			})

		}
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "playlist isn't playing",
		})
	}

}

func (h *Handler) AddSong(c *gin.Context) {
	var input gocloudcamppart2.Track
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.ChangePlaylist.AddSong(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "song has added",
		})
	}
}

func (h *Handler) DeleteSong(c *gin.Context) {
	input := c.Param("id")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if h.services.ChangePlaylist.GetTrackD(id) {
		if playlist.Played && playlist.TrackID == id {
			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "[WARNING]this track is playing right now",
			})
		} else {
			err := h.services.ChangePlaylist.DeleteSong(id)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "track has deleted",
			})

		}
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "track wasn't find",
		})
	}

}

func (h *Handler) StopPlaylist(c *gin.Context) {
	if playlist.Played {
		playlist.Stoped = true
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "playlist has stoped",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "playlist isn't playing",
		})
	}
}

func (h *Handler) ChangeSong(c *gin.Context) {
	input := c.Param("id")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inputTrack gocloudcamppart2.Track
	if err := c.BindJSON(&inputTrack); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputTrack.ID = id

	if playlist.Played && playlist.TrackID == id {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answer": "[WARNING]this track is playing right now",
		})
	} else {
		if h.services.ChangePlaylist.GetTrackD(id) {
			err := h.services.ChangePlaylist.ChangeSong(inputTrack)

			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "track has changed",
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"answer": "track wasn't find",
			})

		}

	}
}

func (h *Handler) ShowSong(c *gin.Context) {
	var tracks []gocloudcamppart2.Track
	tracks, err := h.services.ChangePlaylist.ShowSong()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, track := range tracks {
		c.JSON(http.StatusOK, map[string]interface{}{
			"answ": track,
		})
	}
}
