package main

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/icons"
)

const (
	title = "U-Rule"
	motto = "You are the best. Always have been."
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

	dismissed bool
}

func (b *Body) Render() vecty.ComponentOrHTML {
	spa := &mdc.SPA{
		FullHeightDrawer: false,
		Drawer: &mdc.Leftbar{
			Title:     vecty.Text(title),
			Subtitle:  vecty.Text(motto),
			Variant:   mdc.VariantDismissableLeftbar,
			Dismissed: b.dismissed,
			List: &mdc.List{
				ID:       "list-111",
				ListElem: mdc.ElementNavigationList,
				Items: vecty.List{
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
			b.dismissed = !b.dismissed
			globalListener()
			// spa.Drawer.Dismiss(!spa.Drawer.IsDismissed())
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
