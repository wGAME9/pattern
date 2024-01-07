package command

import "github.com/wGAME9/pattern/command/player"

var _ Command = (*ButtonUpArrow)(nil)

type ButtonUpArrow struct{}

func (c *ButtonUpArrow) Execute(p *player.Player) {
	p.MoveUp()
}

var _ Command = (*ButtonDownArrow)(nil)

type ButtonDownArrow struct{}

func (c *ButtonDownArrow) Execute(p *player.Player) {
	p.MoveDown()
}

var _ Command = (*ButtonLeftArrow)(nil)

type ButtonLeftArrow struct{}

func (c *ButtonLeftArrow) Execute(p *player.Player) {
	p.MoveLeft()
}

var _ Command = (*ButtonRightArrow)(nil)

type ButtonRightArrow struct{}

func (c *ButtonRightArrow) Execute(p *player.Player) {
	p.MoveRight()
}
