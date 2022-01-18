package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
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
		Navbar: &mdc.Navbar{
			// Variant: mdc.VariantTopBarFixed,
			SectionStart: vecty.List{
				&mdc.Typography{Root: vecty.Text(title), Style: mdc.Headline6},
			},
		},
		Drawer: &mdc.Leftbar{
			// Title:       vecty.Text(title),
			// Subtitle:    vecty.Text(motto),
			Dismissible: true,
			List: &mdc.List{
				ListElem: mdc.ElementDivList,
				List: vecty.List{
					&mdc.ListItem{Label: vecty.Text("Visit our hot sale"), Icon: icons.PointOfSale},
					&mdc.ListItem{Label: vecty.Text("Our mission"), Icon: icons.AirplanemodeActive},
					&mdc.ListItem{Label: vecty.Text("Our values"), Icon: icons.PersonalInjury},
					&mdc.ListItem{Label: vecty.Text("Contact us"), Icon: icons.ContactPhone},
				},
			},
		},
		Content: &mdc.Typography{Root: vecty.Text("The future, all in one place. Now. Buy now.")},
	}
	return elem.Body(
		spa,
	)
}
