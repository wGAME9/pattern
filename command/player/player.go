package player

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x        int
	y        int
	radius   int
	velocity int
	color    color.Color
}

func NewPlayer(x, y, velocity int, color color.Color) *Player {
	return &Player{
		x:        x,
		y:        y,
		radius:   10,
		velocity: velocity,
		color:    color,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	radius64 := float64(p.radius)
	minAngle := math.Acos(1 - 1/radius64)

	// Draw a circle as player
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)
		x1 := int(math.Round(float64(p.x) + xDelta))
		y1 := int(math.Round(float64(p.y) + yDelta))
		screen.Set(x1, y1, p.color)
	}
}

func (p *Player) MoveUp() {
	p.y -= p.velocity
}

func (p *Player) MoveDown() {
	p.y += p.velocity
}

func (p *Player) MoveRight() {
	p.x += p.velocity
}

func (p *Player) MoveLeft() {
	p.x -= p.velocity
}
