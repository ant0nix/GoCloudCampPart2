package service

import (
	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/ant0nix/GoCloudCampPart2/pkg/repository"
)

type ChangePlaylistStruct struct {
	repo repository.ChangePlaylist
}

func NewChangePlaylistStruct(repo repository.ChangePlaylist) *ChangePlaylistStruct {
	return &ChangePlaylistStruct{repo: repo}
}

func (s *ChangePlaylistStruct) AddSong(track gocloudcamppart2.Track) error {
	return s.repo.AddSong(track)
}

func (s *ChangePlaylistStruct) DeleteSong(id int) error {

	return s.repo.DeleteSong(id)
}

func (s *ChangePlaylistStruct) ChangeSong(track gocloudcamppart2.Track) error {
	return s.repo.ChangeSong(track)
}

func (s *ChangePlaylistStruct) GetTrackD(id int) bool {
	return s.repo.GetTrackD(id)
}

func (s *ChangePlaylistStruct) ShowSong() ([]gocloudcamppart2.Track, error) {
	return s.repo.ShowSong()
}
