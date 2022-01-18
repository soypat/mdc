package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
	"github.com/soypat/mdc/icons"
)

const (
	title = "U-Rule"
	motto = "You are the best. Always have been."
)

var globalListener func()

func main() {
	jlog.PackageLevel = jlog.LevelTrace
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
	spa := &mdc.SPA{
		FullHeightDrawer: false,
		Drawer: &mdc.Leftbar{
			Title:    vecty.Text(title),
			Subtitle: vecty.Text(motto),
			Variant:  mdc.VariantDismissableLeftbar,
			List: &mdc.List{
				ListElem: mdc.ElementNavigationList,
				List: vecty.List{
					&mdc.ListItem{Label: vecty.Text("Visit our hot sale"), Icon: icons.PointOfSale, ListItemElem: mdc.ElementAnchorListItem, Href: "#dsd"},
					&mdc.ListItem{Label: vecty.Text("Our mission"), Icon: icons.AirplanemodeActive},
					&mdc.ListItem{Label: vecty.Text("Our values"), Icon: icons.PersonalInjury},
					&mdc.ListItem{Label: vecty.Text("Contact us"), Icon: icons.ContactPhone},
				},
			},
		},
		Content: &mdc.Typography{Root: vecty.Text("The future, all in one place. Now. Buy now.")},
	}
	// Dehaze is the three-horizontal-line icon commonly used for
	// indicating a dismissible menu list of action items.
	dehaze := &mdc.Button{
		Icon: icons.Dehaze,
		Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
			spa.Drawer.Dismiss(!spa.Drawer.IsDismissed())
		})},
	}
	spa.Navbar = &mdc.Navbar{
		SectionStart: vecty.List{
			dehaze,
			&mdc.Typography{Root: vecty.Text(title), Style: mdc.Headline6},
		},
	}
	return elem.Body(
		spa,
	)
}
