package repository

import (
	"fmt"
	"log"

	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/jmoiron/sqlx"
)

const (
	insertProcedure = "insert_playlist"
)

type ChangePlaylistRepoStruct struct {
	db *sqlx.DB
}

func NewChangePlaylistRepoStruct(db *sqlx.DB) *ChangePlaylistRepoStruct {
	return &ChangePlaylistRepoStruct{db: db}
}

func (r *ChangePlaylistRepoStruct) AddSong(track gocloudcamppart2.Track) error {
	query := fmt.Sprintf("CALL %s ($1)", insertProcedure)
	_, err := r.db.Exec(query, track.Duration)
	if err != nil {
		log.Printf("Error! repository/AddSong: %s", err.Error())
	}
	return err
}

func (r *ChangePlaylistRepoStruct) DeleteSong(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", playlistTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error! repository/DeleteSong: %s", err.Error())
	}
	return err
}

func (r *ChangePlaylistRepoStruct) ChangeSong(track gocloudcamppart2.Track) error {
	query := fmt.Sprintf("UPDATE %s SET duration = $1 WHERE id = $2", playlistTable)
	_, err := r.db.Exec(query, track.Duration, track.ID)
	if err != nil {
		log.Printf("Error! repository/ChangeSong: %s", err.Error())
	}
	return err
}

func (r *ChangePlaylistRepoStruct) GetTrackD(id int) bool {
	var output []gocloudcamppart2.Track
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", playlistTable)
	err := r.db.Select(&output, query, id)
	if err != nil {
		return false
	}
	if len(output) > 0 {
		return true
	} else {
		return false
	}
}
