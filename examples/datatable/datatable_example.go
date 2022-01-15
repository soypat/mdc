package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

var (
	globalListener func()
	tabData        = []mdc.Series{
		&mdc.StringSeries{
			Label: "Dog breed",
			Data:  []string{"Husky", "Chocolate Labrador", "Beagle"},
		},
		&mdc.StringSeries{
			Label: "Rating",
			Data:  []string{"5/5", "7/5", "6/5"},
		},
	}
	tabRows = 3
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

	disableButton bool
}

func (b *Body) Render() vecty.ComponentOrHTML {
	dt := &mdc.DataTable{
		Columns: tabData,
		Rows:    tabRows,
	}
	return elem.Body(
		dt,
	)
}
