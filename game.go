package main

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen  tcell.Screen
	Ball    Ball
	Player1 Player
	Player2 Player
}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) GameOver() bool {
	return g.Player1.Score == 3 || g.Player2.Score == 3
}

func (g *Game) DeclareWinner() string {
	if !g.GameOver() {
		return "Game Not Over. No Winner"
	}

	if g.Player1.Score > g.Player2.Score {
		return "Left Player"
	} else {
		return "Right Player"
	}
}

func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()

		width, height := s.Size()

		if g.GameOver() {
			drawSprite(s, (width/2)-4, 7, (width/2)+5, 7, defStyle, "Game Over")
			drawSprite(s, (width/2)-8, 11, (width/2)+10, 11, defStyle, g.DeclareWinner()+" Wins!")
			s.Show()
			duration := time.Second * 3
			time.Sleep(duration)
			g.Player1.Score = 0
			g.Player2.Score = 0
		}

		if g.Ball.intersects(g.Player1.Paddle) || g.Ball.intersects(g.Player2.Paddle) {
			g.Ball.reverseX()
			g.Ball.reverseY()
		}

		if g.Ball.X <= 0 {
			g.Player2.Score++
			g.Ball.Reset(width/2, height/2, -1, 1)
		}

		if g.Ball.X >= width {
			g.Player1.Score++
			g.Ball.Reset(width/2, height/2, 1, 1)
		}

		// display score
		drawSprite(s, (width/2)-5, 1, 1, 1, defStyle, strconv.Itoa(g.Player1.Score))
		drawSprite(s, (width/2)+5, 1, 1, 1, defStyle, strconv.Itoa(g.Player2.Score))

		g.Ball.CheckEdges(width, height)
		g.Ball.Update()

		s.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		paddleStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

		drawSprite(s,
			g.Player1.Paddle.X,
			g.Player1.Paddle.Y,
			g.Player1.Paddle.X+g.Player1.Paddle.width,
			g.Player1.Paddle.Y+g.Player1.Paddle.height,
			paddleStyle,
			g.Player1.Paddle.Display())

		drawSprite(s,
			g.Player2.Paddle.X,
			g.Player2.Paddle.Y,
			g.Player2.Paddle.X+g.Player2.Paddle.width,
			g.Player2.Paddle.Y+g.Player2.Paddle.height,
			paddleStyle,
			g.Player2.Paddle.Display())

		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}
