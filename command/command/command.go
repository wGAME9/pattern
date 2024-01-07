package command

import (
	"github.com/wGAME9/pattern/command/player"
)

type Command interface {
	Execute(*player.Player)
}

