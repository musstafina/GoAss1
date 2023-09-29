package observers

import "fmt"

type PlayerObserver struct{
	playerID string
}

func NewPlayerObserver(playerID string) *PlayerObserver {
	return &PlayerObserver{playerID}
}

func (po *PlayerObserver) Update(event string, data interface{}) {
	fmt.Printf("Player %s recieved event: %s, Data: %+v\n", po.playerID, event, data)
}