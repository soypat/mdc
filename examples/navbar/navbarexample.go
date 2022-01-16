package main

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
	"github.com/soypat/mdc/icons"
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
	vecty.SetTitle("Welcome visitor")
	butt := &mdc.Button{
		// Uncomment to use text button
		// Label: vecty.Text("Rerender"),
		Icon: icons.Dehaze,
		Listeners: []*vecty.EventListener{
			event.Click(func(e *vecty.Event) {
				counter++
				listener()
			}),
		},
	}

	bar := &mdc.Navbar{
		SectionStart: vecty.List{
			butt,
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

	return elem.Body(
		// vecty.Markup(vecty.UnsafeHTML(example)),
		bar,
	)
}

const example = `<header class="mdc-top-app-bar">
<div class="mdc-top-app-bar__row">
  <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
	<button class="material-icons mdc-top-app-bar__navigation-icon mdc-icon-button" aria-label="Open navigation menu">menu</button>
	<span class="mdc-top-app-bar__title">Page title</span>
  </section>
  <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end" role="toolbar">
	<button class="material-icons mdc-top-app-bar__action-item mdc-icon-button" aria-label="Favorite">favorite</button>
	<button class="material-icons mdc-top-app-bar__action-item mdc-icon-button" aria-label="Search">search</button>
	<button class="material-icons mdc-top-app-bar__action-item mdc-icon-button" aria-label="Options">more_vert</button>
  </section>
</div>
</header>
<main class="mdc-top-app-bar--fixed-adjust">
App content
</main>`
