package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
	"github.com/soypat/mdc/icons"
)

// Global state.
var (
	globalListener func()
	content        = []struct {
		title   string
		icon    icons.Icon
		content string
	}{
		{title: "On Sale!", icon: icons.Savings, content: "Sorry, no items currently on sale"},
		{title: "Aisle", icon: icons.AddShoppingCart, content: "Stock depleted!"},
		{title: "Checkout", icon: icons.ShoppingCartCheckout, content: "You have no items in your cart"},
	}
)

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()
	jlog.PackageLevel = jlog.LevelTrace
	body := &Body{}
	globalListener = func() {
		vecty.Rerender(body)
	}
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
	selected int
}

func (b *Body) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for _, page := range content {
		items = append(items, &mdc.ListItem{
			Label:        elem.Span(&mdc.Typography{Root: vecty.Text(page.title)}), //vecty.Text(page.title),
			Icon:         page.icon,
			ListItemElem: mdc.ElementAnchorListItem,
			Href:         "#",
			Ripple:       true,
		})
	}

	selectedItem := items[b.selected].(*mdc.ListItem)
	selectedItem.Active = true
	lb := &mdc.Leftbar{
		Variant:     mdc.VariantDismissableLeftbar,
		StartClosed: false,
		Title:       vecty.Text("Welcome user"),
		Subtitle:    vecty.Text("NewAge Groceries welcomes you"),
		List: &mdc.List{
			// ID:            "list-1",
			ClickListener: b.listenDehaze,
			List:          items,
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
		Icon: icons.Dehaze,
		Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
			lb.Dismiss(!lb.IsDismissed())
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

func (b *Body) listenDehaze(idx int, e *vecty.Event) {
	b.selected = idx
	globalListener()
}
