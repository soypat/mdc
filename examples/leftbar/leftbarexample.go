package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
)

var listener func()

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()

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
		Label: vecty.Text("Button"),
	}
	lb := &mdc.Leftbar{}
	jlog.Debug("button disabled:", butt.Disabled)
	return elem.Body(
		lb,
	)
}
