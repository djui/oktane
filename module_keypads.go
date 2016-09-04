package main

import (
	"math/rand"

	termbox "github.com/nsf/termbox-go"
)

var (
	symbols = map[rune]rune{
		'O': 'Ϙ', // '\u03d8' | greek    | Greek Letter Archaic Koppa
		'A': 'Ѧ', // '\u0466' | cyrillic | Cyrillic Capital Letter Little Yus
		'^': 'ƛ', // '\u019b' | latin    | Latin Small Letter Lambda With Stroke
		's': 'Ϟ', // '\u03de' | greek    | Greek Letter Koppa
		'H': 'Ѭ', // '\u046c' | cyrillic | Cyrillic Capital Letter Iotified Big Yus
		'x': 'ϗ', // '\u03d7' | greek    | Greek Kai Symbol
		'D': 'Ͽ', // '\u03ff' | greek    | Greek Capital Reversed Dotted Lunate Sigma Symbol
		'E': 'Ӭ', // '\u04ec' | cyrillic | Cyrillic Capital Letter E With Diaeresis
		'0': 'Ҩ', // '\u04a8' | cyrillic | Cyrillic Capital Letter Abkhasian Ha
		'+': '☆', // '\u2606' | symbol   | White Star
		'?': '¿', // '\u00bf' | latin    | Inverted Question Mark
		'c': '©', // '\u00a9' | latin    | Copyright Sign
		'w': 'Ѽ', // '\u047c' | cyrillic | Cyrillic Capital Letter Omega With Titlo
		'K': 'Җ', // '\u0496' | cyrillic | Cyrillic Capital Letter Zhe With Descender
		'R': 'Ԇ', // '\u0506' | cyrillic | Cyrillic Capital Letter Komi Dzje
		'6': 'б', // '\u0431' | cyrillic | Cyrillic Small Letter Be
		'P': '¶', // '\u00b6' | latin    | Pilcrow Sign
		'b': 'Ѣ', // '\u0462' | cyrillic | Cyrillic Capital Letter Yat
		'u': 'ټ', // '\u067c' | arabic   | Arabic Letter Teh With Ring
		'Y': 'Ψ', // '\u03a8' | greek    | Greek Capital Letter Psi
		'C': 'Ͼ', // '\u03fe' | greek    | Greek Capital Dotted Lunate Sigma Symbol
		'3': 'Ѯ', // '\u046e' | cyrillic | Cyrillic Capital Letter Ksi
		'*': '★', // '\u2605' | symbol   | Black Star
		'#': '҂', // '\u0482' | cyrillic | Cyrillic Thousands Sign
		'a': 'æ', // '\u00e6' | latin    | Latin Small Letter Ae
		'N': 'Ҋ', // '\u048a' | cyrillic | Cyrillic Capital Letter Short I With Tail
		'o': 'Ω', // '\u03a9' | greek    | Greek Capital Letter Omega
	}

	symbolColumns = [][]rune{
		{'O', 'A', '^', 's', 'H', 'x', 'D'},
		{'E', 'O', 'D', '0', '+', 'x', '?'},
		{'c', 'w', '0', 'K', 'R', '^', '+'},
		{'6', 'P', 'b', 'H', 'K', '?', 'u'},
		{'Y', 'u', 'b', 'C', 'P', '3', '*'},
		{'6', 'E', '#', 'a', 'Y', 'N', 'o'},
	}
)

type KeypadsModule struct {
	BasicModule

	solved bool
}

func (m *KeypadsModule) solve(b bool) {
	m.solved = b
}

func (m *KeypadsModule) render(r Rect, c Config) {
	m.renderFrame(r, c)
	m.renderIndicator(r, c)
	m.renderKeypads(r, c)
}

func (m *KeypadsModule) renderKeypads(r Rect, c Config) {
	renderRectThin(Rect{Pos{r.x + 3, r.y + 1}, Size{5, 3}})
	renderRectThin(Rect{Pos{r.x + r.w - 9, r.y + 1}, Size{5, 3}})
	renderRectThin(Rect{Pos{r.x + 3, r.y + r.h - 4}, Size{5, 3}})
	renderRectThin(Rect{Pos{r.x + r.w - 9, r.y + r.h - 4}, Size{5, 3}})

	color := termbox.ColorGreen
	if c.noColor {
		color = termbox.ColorDefault
	}

	termbox.SetCell(r.x+3+2, r.y+1, '┰', color, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w-9+2, r.y+r.h-4, '┰', color, termbox.ColorDefault)

	col := symbolColumns[rand.Intn(len(symbolColumns))]
	idx := rand.Perm(len(col))
	s := []rune{
		symbols[col[idx[0]]],
		symbols[col[idx[1]]],
		symbols[col[idx[2]]],
		symbols[col[idx[3]]],
	}
	if c.noANSI {
		s = []rune{
			col[idx[0]],
			col[idx[1]],
			col[idx[2]],
			col[idx[3]],
		}
	}
	termbox.SetCell(r.x+3+2, r.y+2, s[0], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w-9+2, r.y+2, s[1], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+3+2, r.y+r.h-4+1, s[2], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(r.x+r.w-9+2, r.y+r.h-4+1, s[3], termbox.ColorDefault, termbox.ColorDefault)

}
