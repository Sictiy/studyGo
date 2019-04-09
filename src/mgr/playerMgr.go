package mgr

import (
	"awesomeProject/src/entry/interface"
	"sync"
)

var pPlayerMgr *playerMgr
var oncePlayer sync.Once

func GetPlayerMgr() *playerMgr {
	oncePlayer.Do(func() {
		pPlayerMgr = &playerMgr{}
	})
	return pPlayerMgr
}

type playerMgr struct {
	playersList []*_interface.IPlayer
}

func (pPlayerMgr *playerMgr) StartDoAction() {
	for _, player := range pPlayerMgr.playersList {
		(*player).StartDoAction()
	}
}

func (pPlayerMgr *playerMgr) AddPlayers(players [] *_interface.IPlayer) {
	for i := 0; i < len(players) ; i++ {
		pPlayerMgr.playersList = append(pPlayerMgr.playersList, players[i])
	}
}

func (pPlayerMgr *playerMgr) AddPlayer(player *_interface.IPlayer) {
	pPlayerMgr.playersList = append(pPlayerMgr.playersList, player)
}