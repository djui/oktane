package main

import termbox "github.com/nsf/termbox-go"

type KnobModule struct {
	BasicModule
}

func (m *KnobModule) solve(b bool) {
}

func (m *KnobModule) render(r Rect, c Config) {
	m.renderFrame(r, c)

	termbox.SetCell(r.x+4+0, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+0, 'K', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+0, 'N', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+0, 'O', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+0, 'B', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+0, ' ', termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(r.x+4+0, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
}
