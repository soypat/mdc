package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

const (
	title = "U-Rule"
	motto = "You are the best. Always have been."
)

var globalListener func()

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()

	body := &Body{}
	globalListener = func() {
		vecty.Rerender(body)
	}
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
}

func (b *Body) Render() vecty.ComponentOrHTML {
	navbar := &mdc.Navbar{
		SectionStart: vecty.List{
			&mdc.Typography{Root: vecty.Text(title), Style: mdc.Headline6},
		},
	}
	return elem.Body(
		navbar,
	)
}
