package main

import termbox "github.com/nsf/termbox-go"

type SimonSaysModule struct {
	BasicModule
}

func (m *SimonSaysModule) solve(b bool) {
}

func (m *SimonSaysModule) render(r Rect, c Config) {
	m.renderFrame(r, c)

	termbox.SetCell(r.x+4+0, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+0, 'S', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+0, 'I', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+0, 'M', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+0, 'O', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+0, 'N', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(r.x+4+0, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+1, 'S', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+1, 'A', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+1, 'Y', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+1, 'S', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
}
