package main

import (
	"project/game_mode"
	"project/observers"
	"project/players"
	"project/utils"
)
func main() {
	eventDispatcher := utils.NewEventDispatcher()

	playerManager := players.NewPlayerManager()
	player1 :=	players.NewPlayer("player1")
	player2 := players.NewPlayer("player2")
	playerManager.AddPlayer(player1)
	playerManager.AddPlayer(player2)

	playerObserver1 := observers.NewPlayerObserver(player1.ID)
	playerObserver2 := observers.NewPlayerObserver(player2.ID)
	eventDispatcher.Register("game_event", playerObserver1)
	eventDispatcher.Register("game_event", playerObserver2)

	context := &game_mode.GameContext{}

	singleplayerMode := &game_mode.SinglePlayer{}
	multiplayerMode := &game_mode.Multiplayer{}

	context.SetMode(singleplayerMode)
	context.Play()

	eventDispatcher.Dispatch("game_event", "Game Over")

	context.SetMode(multiplayerMode)
	context.Play()

	eventDispatcher.Dispatch("game_event", "Game Over")


}