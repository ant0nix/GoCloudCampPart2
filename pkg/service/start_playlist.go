package service

import (
	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/ant0nix/GoCloudCampPart2/pkg/repository"
)

type StartPlaylistStruct struct {
	repo repository.StartPLaylist
}

func NewStartPlaylistStruct(repo repository.StartPLaylist) *StartPlaylistStruct {
	return &StartPlaylistStruct{repo: repo}
}

func (r *StartPlaylistStruct) NextSong(id int) (gocloudcamppart2.Track, error) {
	track, err := r.repo.GetTrack(id)
	if err != nil {
		return track, err
	}
	if track.ID == 0 {
		return gocloudcamppart2.Track{ID: -1}, nil
	} else {

		return track, nil
	}
}

func (r *StartPlaylistStruct) LastSong() int {
	return r.repo.LastSong()
}

func (r *StartPlaylistStruct) FirstSong() int {
	return r.repo.FirstSong()
}
