package game_mode

import (
	"fmt"
)

type SinglePlayer struct {}

func (sp *SinglePlayer) Start() {
	fmt.Println("Single-player game started")
}

func (sp *SinglePlayer) End() {
	fmt.Println("Single-player game ended")
}