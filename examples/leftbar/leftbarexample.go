package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
)

// Global state.
var (
	globalListener func()
	content        = []struct {
		title   string
		icon    mdc.IconType
		content string
	}{
		{title: "On Sale!", icon: mdc.IconSavings, content: "Sorry, no items currently on sale"},
		{title: "Aisle", icon: mdc.IconAddShoppingCart, content: "Stock depleted!"},
		{title: "Checkout", icon: mdc.IconShoppingCartCheckout, content: "You have no items in your cart"},
	}
)

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

	barOpen  bool
	selected int
}

func (b *Body) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for _, page := range content {
		items = append(items, &mdc.ListItem{
			Label: vecty.Text(page.title),
			Icon:  page.icon,
		})
	}

	selectedItem := items[b.selected].(*mdc.ListItem)
	selectedItem.Active = true

	lb := &mdc.Leftbar{
		Dismissible: true,
		Closed:      !b.barOpen,
		Title:       vecty.Text("Welcome user"),
		Subtitle:    vecty.Text("NewAge Groceries welcomes you"),
		List: &mdc.List{
			ClickListener: func(idx int, e *vecty.Event) {
				b.selected = idx
				globalListener()
			},
			List: items,
		},
	}

	// TODO(soypat): Add material animation
	// From material.io, the js looks like this:
	//
	// import {MDCTopAppBar} from "@material/top-app-bar";
	// const topAppBar = MDCTopAppBar.attachTo(document.getElementById('app-bar'));
	// topAppBar.setScrollTarget(documen ,t.getElementById('main-content'));
	// topAppBar.listen('MDCTopAppBar:nav', () => {
	// drawer.open = !drawer.open;
	// });
	but := &mdc.Button{
		Icon: mdc.IconDehaze,
		Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
			b.barOpen = !b.barOpen
			globalListener()
		})},
	}
	return elem.Body(
		lb,
		elem.Div(vecty.Markup(vecty.Class("main-content", "mdc-drawer-app-content")),
			elem.Main(
				but,
				&mdc.Typography{
					Root: vecty.Text(content[b.selected].content),
				},
				&mdc.Checkbox{
					ID:    "checkbox-1",
					Label: vecty.Text("I acknowledge this fact."),
				},
			),
		),
	)
}
