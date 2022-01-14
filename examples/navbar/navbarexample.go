package main

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
)

// Global state
var (
	counter  = 1
	listener func()
)

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()
	jlog.PackageLevel = jlog.LevelTrace
	body := &Body{}
	listener = func() {
		vecty.Rerender(body)
	}
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
}

func (b *Body) Render() vecty.ComponentOrHTML {
	butt := &mdc.Button{
		Label: vecty.Text("Rerender"),
	}

	bar := &mdc.Navbar{
		SectionStart: vecty.List{
			elem.Div(
				vecty.Markup(event.Click(func(e *vecty.Event) {
					counter++
					listener()
				})),
				butt,
			),
		},
		SectionCenter: vecty.List{
			&mdc.Typography{
				Root:  vecty.Text("soypat's mancave"),
				Style: mdc.Headline6,
			},
		},
		SectionEnd: vecty.List{
			&mdc.Typography{
				Root:  vecty.Text("you are visitor #" + strconv.Itoa(counter)),
				Style: mdc.Headline5,
			},
		},
	}

	return elem.Body(bar)
}
