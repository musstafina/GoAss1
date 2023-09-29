package game_mode

type GameMode interface {
	Start()
	End()
}

type GameContext struct {
	mode GameMode
}

func (gc *GameContext) SetMode(mode GameMode) {
	gc.mode = mode
}

func (gc *GameContext) Play() {
	gc.mode.Start()
	gc.mode.End()
}