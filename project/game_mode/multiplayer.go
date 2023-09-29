package game_mode

import (
	"fmt"
)

type Multiplayer struct{}

func (mp *Multiplayer) Start() {
	fmt.Println("Multiplayer game started")
}

func (mp *Multiplayer) End() {
	fmt.Println("Multiplayer game ended")
}