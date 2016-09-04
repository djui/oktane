package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

const (
	digit0 = 490 // 10101111
	digit1 = 288 // 00001001
	digit2 = 242 // 10011110
	digit3 = 434 // 10011011
	digit4 = 312 // 00111001
	digit5 = 410 // 10110011
	digit6 = 474 // 10110111
	digit7 = 290 // 10001001
	digit8 = 507 // 10111111
	digit9 = 442 // 10111011
)

var (
	digit  = []rune{' ', '_', ' ', '|', ' ', '|', '|', '_', '|'}
	digits = []int{digit0, digit1, digit2, digit3, digit4, digit5, digit6, digit7, digit8, digit9}
)

type TimeModule struct {
	BasicModule

	solved  bool
	strikes int
	time    time.Duration
}

func (m *TimeModule) solve(b bool) {
	m.solved = b
}

func (m *TimeModule) render(r Rect, c Config) {
	m.strikes = 2
	m.renderFrame(r, c)
	m.renderStrikes(r, c)
	m.renderTime(r, c)
}

func (m *TimeModule) renderStrikes(r Rect, c Config) {
	renderRectThin(Rect{Pos{r.x + 5, r.y + 1}, Size{r.w - (2 * 5), 3}})

	color := termbox.ColorRed
	if c.noColor {
		color = termbox.ColorDefault
	}
	color |= termbox.AttrBold

	if m.strikes >= 1 {
		termbox.SetCell(r.x+7, r.y+2, 'x', color, termbox.ColorDefault)
	}
	if m.strikes >= 2 {
		termbox.SetCell(r.x+r.w-7-1, r.y+2, 'x', color, termbox.ColorDefault)
	}
}

func durationToDigits(d time.Duration) []int {
	ds := []int{}
	ds = append(ds, digits[int(d.Minutes())/10])
	ds = append(ds, digits[int(d.Minutes())%10])
	ds = append(ds, 0)
	ds = append(ds, digits[(int(d.Seconds())%60)/10])
	ds = append(ds, digits[int(d.Seconds())%10])
	return ds
}

var start = time.Now()

func (m *TimeModule) renderTime(r Rect, c Config) {
	termbox.SetCell(r.x+2*3+2, r.y+5, '.', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+2*3+2, r.y+6, '.', termbox.ColorDefault, termbox.ColorDefault)

	digits := durationToDigits(time.Now().Sub(start))

	for i, d := range digits {
		if d&(1<<1) != 0 {
			termbox.SetCell(r.x+i*3+2, r.y+4+0, '_', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<3) != 0 {
			termbox.SetCell(r.x+i*3+1, r.y+4+1, '|', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<4) != 0 {
			termbox.SetCell(r.x+i*3+2, r.y+4+1, '_', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<5) != 0 {
			termbox.SetCell(r.x+i*3+3, r.y+4+1, '|', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<6) != 0 {
			termbox.SetCell(r.x+i*3+1, r.y+4+2, '|', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<7) != 0 {
			termbox.SetCell(r.x+i*3+2, r.y+4+2, '_', termbox.ColorDefault, termbox.ColorDefault)
		}
		if d&(1<<8) != 0 {
			termbox.SetCell(r.x+i*3+3, r.y+4+2, '|', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
