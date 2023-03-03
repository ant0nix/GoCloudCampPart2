package service

import (
	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/ant0nix/GoCloudCampPart2/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ChangePlaylist interface {
	AddSong(track gocloudcamppart2.Track) error
	DeleteSong(id int) error
	ChangeSong(track gocloudcamppart2.Track) error
	GetTrackD(id int) bool
	ShowSong() ([]gocloudcamppart2.Track, error)
}

type StartPLaylist interface {
	NextSong(id int) (gocloudcamppart2.Track, error)
	LastSong() int
	FirstSong() int
}

type Service struct {
	ChangePlaylist
	StartPLaylist
}

func NewServies(repo *repository.Repository) *Service {
	return &Service{
		ChangePlaylist: NewChangePlaylistStruct(repo),
		StartPLaylist:  NewStartPlaylistStruct(repo),
	}
}
