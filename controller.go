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

	// Create channel, that can communicate accross threads
	eventQueue := make(chan termbox.Event)

	// Function is handling termbox events such as user input
	go func() {
		// Infinite loop
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := NewGame()
	render(g)

	for {
		// switch for channels
		select {
		case ev := <-eventQueue:
			// check user input
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q':
					return
				case ev.Key == termbox.KeyArrowLeft:
					g.movePiece(left)
				case ev.Key == termbox.KeyArrowRight:
					g.movePiece(right)
				}

			}

		//case <-g.clock.C:
		//timer went of
		//g.play()
		default:
			render(g)
			time.Sleep(animationSpeed)
		}

	}

}
