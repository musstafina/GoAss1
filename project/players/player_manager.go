package players

type PlayerManager struct{
	players map[string]*Player
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		players: make(map[string]*Player),
	}
}

func (pm *PlayerManager) AddPlayer(player *Player) {
	pm.players[player.ID] = player
}

func (pm *PlayerManager) GetPlayer(playerID string) *Player {
	return pm.players[playerID]
}
