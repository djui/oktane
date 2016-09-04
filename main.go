package main

import (
	"log"
	"time"

	termbox "github.com/nsf/termbox-go"
)

const animationSpeed = 10 * time.Millisecond

func main() {
	err := termbox.Init()
	assertNoError(err, "termbox init failed")
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	c := Config{
		noColor: false,
		noANSI:  false,
	}
	g := NewGame(c)
	render(g)

	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowUp:
				//g.move(UP)
			case ev.Key == termbox.KeyArrowDown:
				//g.move(DOWN)
			case ev.Key == termbox.KeyArrowLeft:
				//g.move(LEFT)
			case ev.Key == termbox.KeyArrowRight:
				//g.move(RIGHT)
			case ev.Key == termbox.KeyEsc:
				return
			}
		}
		render(g)
		time.Sleep(animationSpeed)
	}
}

func assertNoError(err error, msg string) {
	if err != nil {
		log.Panicf("Error: %s: %v", msg, err)
	}
}
