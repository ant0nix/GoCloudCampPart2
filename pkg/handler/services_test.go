package handler

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"

	gocloudcamppart2 "github.com/ant0nix/GoCloudCampPart2"
	"github.com/ant0nix/GoCloudCampPart2/pkg/service"
	mock_service "github.com/ant0nix/GoCloudCampPart2/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

var (
	badRequest = `{"message":"json: cannot unmarshal string into Go struct field Track.duration of type int"}`
)

func TestHandler_AddSong(t *testing.T) {

	type mockBehavior func(s *mock_service.MockChangePlaylist, track gocloudcamppart2.Track)

	testTable := []struct {
		name            string
		inputBody       string
		inputTrack      gocloudcamppart2.Track
		mockBehavior    mockBehavior
		expectedStatus  int
		expectedRequest string
	}{
		{
			name:      "OK",
			inputBody: `{"duration": 6}`,
			inputTrack: gocloudcamppart2.Track{
				Duration: 6,
			},
			mockBehavior: func(s *mock_service.MockChangePlaylist, track gocloudcamppart2.Track) {
				s.EXPECT().AddSong(track).Return(nil)
			},
			expectedStatus:  200,
			expectedRequest: `{"answer":"song has added"}`,
		},
		{
			name:            "Bad Request",
			inputBody:       `{"duration": "abc"}`,
			mockBehavior:    func(s *mock_service.MockChangePlaylist, track gocloudcamppart2.Track) {},
			expectedStatus:  400,
			expectedRequest: badRequest,
		},
	}
	for _, testestCase := range testTable {
		t.Run(testestCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			serv_test := mock_service.NewMockChangePlaylist(c)
			testestCase.mockBehavior(serv_test, testestCase.inputTrack)

			services := &service.Service{ChangePlaylist: serv_test}

			handler := NewHandler(services)

			r := gin.New()
			r.POST("/add-song", handler.AddSong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add-song", bytes.NewBufferString(testestCase.inputBody))

			r.ServeHTTP(w, req)
			log.Printf("CODE:%d\nBODY:%s", w.Code, w.Body.String())
			assert.Equal(t, testestCase.expectedStatus, w.Code)
			assert.Equal(t, testestCase.expectedRequest, w.Body.String())
		})
	}
}