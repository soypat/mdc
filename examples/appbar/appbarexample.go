package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
)

var listener func()

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
		// Icon:  mdc.IconBookmark,
	}
	_ = butt
	bar := &mdc.Navbar{
		SectionStart: vecty.List{
			elem.Div(
				butt.SetEventListeners(event.Click(func(e *vecty.Event) {
					listener()
				})),
			),
		},
		// SectionCenter: vecty.List{
		// 	&mdc.Typography{
		// 		Style: mdc.Headline6,
		// 		Root:  vecty.Text("soypat's stuff"),
		// 	},
		// },
	}
	jlog.Debugf("%+v", bar)
	jlog.Debug(bar.SectionStart)
	return elem.Body(bar)
}
