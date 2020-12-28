package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const animationSpeed = 10 * time.Millisecond

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Note: Create channel, that can communicate accross threads
	eventQueue := make(chan termbox.Event)

	// func is handling termbox events such as user input
	go func() {
		// Note: Infinite loop
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := NewGame()
	render(g)

	for {
		// Note: Select is a switch for channels
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q':
					return
				case ev.Key == termbox.KeyArrowDown:
					g.movePiece(down)
				case ev.Key == termbox.KeyArrowUp:
					g.rotatePiece()
				case ev.Key == termbox.KeyArrowLeft:
					g.movePiece(left)
				case ev.Key == termbox.KeyArrowRight:
					g.movePiece(right)
				}
			}

		case <-g.clock.C:
			g.play()
		default:
			render(g)
			time.Sleep(animationSpeed)
		}
	}
}
