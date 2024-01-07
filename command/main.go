package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wGAME9/pattern/command/command"
	"github.com/wGAME9/pattern/command/player"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

// 遊戲
type Game struct {
	player       *player.Player
	inputHandler *InputHandler
}

func (g *Game) Update() error {
	g.inputHandler.HandleInput(g.player)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Command Pattern Example")

	startdAtCenterX := screenWidth / 2
	startedAtCenterY := screenHeight / 2
	playerMovementVelocity := 2
	colorPurple := color.RGBA{255, 0, 255, 255}

	player := player.NewPlayer(
		startdAtCenterX,
		startedAtCenterY,
		playerMovementVelocity,
		colorPurple,
	)

	game := &Game{
		player:       player,
		inputHandler: &InputHandler{},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type InputHandler struct{}

func (h *InputHandler) HandleInput(p *player.Player) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveUpCmd := command.ButtonUpArrow{}
		moveUpCmd.Execute(p)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveDownCmd := command.ButtonDownArrow{}
		moveDownCmd.Execute(p)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveLeftCmd := command.ButtonLeftArrow{}
		moveLeftCmd.Execute(p)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveRightCmd := command.ButtonRightArrow{}
		moveRightCmd.Execute(p)
	}
}
