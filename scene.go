package main

import termbox "github.com/nsf/termbox-go"

type Scene struct {
}

func (s *Scene) flush() {
	termbox.Flush()
}

func (s *Scene) clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (s *Scene) renderModule(m Module, r Rect, c Config) {
	m.render(r, c)
}
