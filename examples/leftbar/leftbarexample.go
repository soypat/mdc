package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
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

	barOpen bool
}

func (b *Body) Render() vecty.ComponentOrHTML {
	lb := &mdc.Leftbar{
		Dismissible: true,
		Closed:      !b.barOpen,
		Title:       vecty.Text("Welcome user"),
		Subtitle:    vecty.Text("NewAge Groceries welcomes you"),
		List: &mdc.List{
			List: vecty.List{
				&mdc.ListItem{
					Label: vecty.Text("On Sale!"),
					Icon:  mdc.IconSavings,
				},
				&mdc.ListItem{
					Label:  vecty.Text("Aisle"),
					Icon:   mdc.IconAddShoppingCart,
					Active: true,
				},
				&mdc.ListItem{
					Label: vecty.Text("Checkout"),
					Icon:  mdc.IconShoppingCartCheckout,
				},
			},
		},
	}
	// TODO(soypat): Add material animation
	// From material.io, the js looks like this:
	//
	// import {MDCTopAppBar} from "@material/top-app-bar";
	// const topAppBar = MDCTopAppBar.attachTo(document.getElementById('app-bar'));
	// topAppBar.setScrollTarget(document.getElementById('main-content'));
	// topAppBar.listen('MDCTopAppBar:nav', () => {
	// drawer.open = !drawer.open;
	// });
	but := &mdc.Button{
		Icon: mdc.IconDehaze,
		Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
			b.barOpen = !b.barOpen
			listener()
		})},
	}
	return elem.Body(
		lb,
		elem.Div(vecty.Markup(vecty.Class("main-content", "mdc-drawer-app-content")),
			elem.Main(
				but,
			),
		),
	)
}
