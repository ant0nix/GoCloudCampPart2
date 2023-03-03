package repository

import (
	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/jmoiron/sqlx"
)

type ChangePlaylist interface {
	AddSong(track gocloudcamppart2.Track) error
	DeleteSong(id int) error
	ChangeSong(track gocloudcamppart2.Track) error
	GetTrackD(id int) bool
}

type StartPLaylist interface {
	GetTrack(id int) (gocloudcamppart2.Track, error)
	LastSong() int
	FirstSong() int
}

type Repository struct {
	ChangePlaylist
	StartPLaylist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ChangePlaylist: NewChangePlaylistRepoStruct(db),
		StartPLaylist:  NewStartPlaylistRepoStruct(db),
	}
}
