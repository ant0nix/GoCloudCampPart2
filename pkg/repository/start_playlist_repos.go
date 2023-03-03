package repository

import (
	"fmt"
	"log"

	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/jmoiron/sqlx"
)

const (
	playlistTable = "playlist"
)

type StartPlaylistRepoStruct struct {
	db *sqlx.DB
}

func NewStartPlaylistRepoStruct(db *sqlx.DB) *StartPlaylistRepoStruct {
	return &StartPlaylistRepoStruct{db: db}
}

func (r *StartPlaylistRepoStruct) GetTrack(id int) (gocloudcamppart2.Track, error) {
	var output []gocloudcamppart2.Track
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", playlistTable)
	err := r.db.Select(&output, query, id)
	if err != nil {
		return gocloudcamppart2.Track{}, err
	}
	if len(output) > 0 {
		return output[0], nil
	} else {
		return gocloudcamppart2.Track{}, nil
	}
}
func (r *StartPlaylistRepoStruct) LastSong() int {
	var output []gocloudcamppart2.Track
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT 1;", playlistTable)
	err := r.db.Select(&output, query)
	if err != nil {
		log.Printf("Error! repository/LastSong: %s", err.Error())
		return 0
	}
	if len(output) > 0 {

		return output[0].ID
	} else {
		return 0
	}

}

func (r *StartPlaylistRepoStruct) FirstSong() int {
	var output []gocloudcamppart2.Track
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC LIMIT 1;", playlistTable)
	err := r.db.Select(&output, query)
	if err != nil {
		log.Printf("Error! repository/FirstSong: %s", err.Error())
		return 0
	}
	if len(output) > 0 {

		return output[0].ID
	} else {
		return 0
	}
}
