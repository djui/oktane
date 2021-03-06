package main

import termbox "github.com/nsf/termbox-go"

type CapacitorDischargeModule struct {
	BasicModule
}

func (m *CapacitorDischargeModule) solve(b bool) {
}

func (m *CapacitorDischargeModule) render(r Rect, c Config) {
	m.renderFrame(r, c)

	termbox.SetCell(r.x+4+0, r.y+3+0, 'C', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+0, 'A', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+0, 'P', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+0, 'A', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+0, 'C', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+0, 'I', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+0, 'T', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+0, 'O', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+0, 'R', termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(r.x+4+0, r.y+3+1, 'D', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+1, r.y+3+1, 'I', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+2, r.y+3+1, 'S', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+3, r.y+3+1, 'C', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+4, r.y+3+1, 'H', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+5, r.y+3+1, 'A', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+6, r.y+3+1, 'R', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+7, r.y+3+1, 'G', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+4+8, r.y+3+1, 'E', termbox.ColorDefault, termbox.ColorDefault)
}
