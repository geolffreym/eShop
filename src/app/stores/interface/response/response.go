package stres

import (
	"app/stores/interface/store"
)

type Response struct {
	Results []store.Product
	Data    map[string]RDetails
}


type RThread struct {
	Items        []store.Product
	TotalResults int
	TotalPages   int
	Store        string
	ChannelSize  int
}

type RDetails struct {
	Total     int
	Pages     int
	Timestamp int64
}

