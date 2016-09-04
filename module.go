package main

import termbox "github.com/nsf/termbox-go"

type Module interface {
	render(r Rect, c Config)
	solve(b bool)
}

type Pos struct {
	x, y int
}

type Size struct {
	w, h int
}

type Rect struct {
	Pos
	Size
}

type BasicModule struct {
}

func (m *BasicModule) renderFrame(r Rect, c Config) {
	renderRectThick(r)
}

func (m *KeypadsModule) renderIndicator(r Rect, c Config) {
	indicator := '\u25CB'
	if c.noANSI {
		indicator = 'o'
	}
	color := termbox.ColorDefault

	if m.solved {
		indicator = '\u25CF'
		if c.noANSI {
			indicator = '*'
		}

		color = termbox.ColorGreen
		if c.noColor {
			color = termbox.ColorDefault
		}
		color |= termbox.AttrBold
	}

	termbox.SetCell(r.x+r.w-2-1, r.y+1, indicator, color, termbox.ColorDefault)
}

func renderRectThick(r Rect) {
	r.w--
	r.h--
	termbox.SetCell(r.x, r.y, '╔', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w, r.y, '╗', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x, r.y+r.h, '╚', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w, r.y+r.h, '╝', termbox.ColorWhite, termbox.ColorDefault)

	for i := 1; i < r.w; i++ {
		termbox.SetCell(r.x+i, r.y, '═', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(r.x+i, r.y+r.h, '═', termbox.ColorWhite, termbox.ColorDefault)
	}
	for i := 1; i < r.h; i++ {
		termbox.SetCell(r.x, r.y+i, '║', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(r.x+r.w, r.y+i, '║', termbox.ColorWhite, termbox.ColorDefault)
	}
}

func renderRectThin(r Rect) {
	r.w--
	r.h--
	termbox.SetCell(r.x, r.y, '┌', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w, r.y, '┐', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x, r.y+r.h, '└', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w, r.y+r.h, '┘', termbox.ColorWhite, termbox.ColorDefault)

	for i := 1; i < r.w; i++ {
		termbox.SetCell(r.x+i, r.y, '─', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(r.x+i, r.y+r.h, '─', termbox.ColorWhite, termbox.ColorDefault)
	}
	for i := 1; i < r.h; i++ {
		termbox.SetCell(r.x, r.y+i, '│', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(r.x+r.w, r.y+i, '│', termbox.ColorWhite, termbox.ColorDefault)
	}
}
