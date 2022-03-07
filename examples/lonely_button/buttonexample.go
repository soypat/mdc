package main

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
)

var globalListener func()

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts(500 * time.Millisecond)

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
	return elem.Body(
		tooltip,
		elem.Main(
			elem.Div(
				butt.SetEventListeners(event.Click(func(e *vecty.Event) {
					b.disableButton = true
					go func() {
						// After one second reenable button.
						time.Sleep(time.Second)
						b.disableButton = false
						globalListener()
					}()
					// Best practices in Vecty are to do top-down renders.
					// See `todomvc` example over at https://github.com/hexops/vecty/tree/main/example
					globalListener()
				})),
			),
		),
	)
}
