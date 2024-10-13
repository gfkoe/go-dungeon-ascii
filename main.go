package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

type Position struct {
	X, Y int
}

var player = Position{5, 5}

func main() {

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer screen.Fini()

	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(style)

	draw(screen)
  for {
    ev := screen.PollEvent()
    switch ev := ev.(type) {
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC:
        return
      case tcell.KeyUp :
        player.Y--
      case tcell.KeyDown:
        player.Y++
      case tcell.KeyLeft:
        player.X--
      case tcell.KeyRight:
        player.X++
      }
      draw(screen)
    case *tcell.EventResize:
      screen.Sync()
    }
  }

}

func draw(screen tcell.Screen) {
	screen.Clear()
	drawRune(screen, player.X, player.Y, '@', tcell.ColorGreen)
	screen.Show()
}

func drawRune(screen tcell.Screen, x, y int, r rune, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color)
	screen.SetContent(x, y, r, nil, style)

}
