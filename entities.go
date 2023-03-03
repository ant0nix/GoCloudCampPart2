package gocloudcamppart2

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Track struct {
	ID       int `json:"-"`
	Duration int `json:"duration" binding:"required"`
}

type Playlist struct {
	FirstId   int
	TrackID   int
	LastTrack int
	Paused    bool
	Played    bool
	NextTrack bool
	PrevTrack bool
	Stoped    bool
	Cond      *sync.Cond
}

func NewPlaylist() *Playlist {
	return &Playlist{
		TrackID:   0,
		Cond:      sync.NewCond(&sync.Mutex{}),
		LastTrack: 0,
	}
}

func (p *Playlist) Play(opChan chan string, track Track) {
	msg := fmt.Sprintf("now play track:%d Duration:%d", track.ID, track.Duration)
	opChan <- msg
	for index := 0; index < track.Duration; index++ {
		if p.NextTrack {
			p.NextTrack = false
			p.TrackID++
			msg = "next track >>>"
			opChan <- msg
			close(opChan)
			return
		}
		if p.PrevTrack {
			if p.TrackID == p.FirstId {
				log.Println("It's a fist track in playlist")
				p.PrevTrack = false
				msg = "[WARNING!] It's a fist track in playlist"
				opChan <- msg
			} else {
				msg = "<<< prev track"
				opChan <- msg
				close(opChan)
				return
			}
		}
		p.Cond.L.Lock()
		if p.Paused {
			msg = "|| paused track"
			opChan <- msg

		}
		if p.Stoped {
			msg = "stop playlist"
			opChan <- msg
			close(opChan)
			return
		}
		var flag bool
		for p.Paused {
			flag = true
			p.Cond.Wait()
		}
		if flag {
			msg = "|> resume track"
			opChan <- msg
		}
		p.Cond.L.Unlock()
		time.Sleep(time.Second)
	}
	close(opChan)
}
