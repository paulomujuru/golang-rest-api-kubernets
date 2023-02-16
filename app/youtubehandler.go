package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type YoutubeHandler struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channel_name"`
	MinutesWatched int    `json:"minutes_watched"`
	VideosWatched  int    `json:"videos_watched"`
}

func getChannelStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		yt := YoutubeHandler{
			Subscribers:    100,
			ChannelName:    "My Channel",
			MinutesWatched: 1000,
			VideosWatched:  100,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(yt); err != nil {
			panic(err)
		}
	}
}
