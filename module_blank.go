package main

import termbox "github.com/nsf/termbox-go"

type BlankModule struct {
	BasicModule
}

func (m *BlankModule) solve(b bool) {
}

func (m *BlankModule) render(r Rect, c Config) {
	m.renderFrame(r, c)
	m.renderPattern(r, c)
}

func (m *BlankModule) renderPattern(r Rect, c Config) {
	for x := 1; x < r.w-1; x++ {
		if x%2 == 0 {
			for y := 1; y < r.h-1; y++ {
				termbox.SetCell(r.x+x, r.y+y, '|', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}
