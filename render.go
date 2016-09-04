package main

type Renderer interface {
	render(x, y int, c Config)
}

const mW = 17
const mH = 8

func render(g *Game) {
	if g.config.noColor {
		// TODO(uwe): Not sure if that's what I want
		// termbox.SetOutputMode(termbox.OutputGrayscale)
	}

	scene := Scene{}
	scene.clear()
	renderRectThin(Rect{Pos{0, 0}, Size{80, 25}})
	for i, m := range g.modules {
		if i >= 6 {
			break
		}
		col := i % 3
		row := i / 3
		r := Rect{Pos{1 + col*(mW+2), 1 + row*(mH+1)}, Size{mW, mH}}
		scene.renderModule(m, r, g.config)

	}
	scene.flush()
}
