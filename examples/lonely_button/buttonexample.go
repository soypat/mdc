package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
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

	disableButton bool
}

func (b *Body) Render() vecty.ComponentOrHTML {
	tooltip := &mdc.Tooltip{
		ID:    "tt-1",
		Label: vecty.Text("This is the button tooltip"),
	}
	butt := &mdc.Button{
		Label: elem.Span(
			vecty.Markup(tooltip),
			vecty.Text("This is button"),
		),
		Disabled: b.disableButton,
	}
	jlog.Debug("button disabled:", butt.Disabled)
	return elem.Body(
		tooltip,
		elem.Main(
			elem.Div(
				butt.SetEventListeners(event.Click(func(e *vecty.Event) {
					jlog.Debug("got a button click!")
					b.disableButton = true
					// Best practices in Vecty are to do top-down renders.
					// See `todomvc` example over at https://github.com/hexops/vecty/tree/main/example
					globalListener()
				})),
			),
		),
	)
}
