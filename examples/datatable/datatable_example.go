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
		&mdc.IntSeries{
			Label: "Average weight",
			Data:  []int{20, 30, 10},
		},
	}
	tabRows = 3
)

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()

	body := &Body{}
	globalListener = func() {
		vecty.Rerender(body)
	}
	vecty.RenderInto("body", body)
	mdc.AddDefaultScripts()
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
}

func (b *Body) Render() vecty.ComponentOrHTML {
	dt := &mdc.DataTable{
		Columns: tabData,
		Rows:    tabRows,
	}
	return elem.Body(
		elem.Main(
			dt,
			elem.Form(
				vecty.Markup(vecty.UnsafeHTML(`<div class="mdc-slider">
				<input class="mdc-slider__input" type="range" min="0" max="100" value="50" name="volume" aria-label="Continuous slider demo">
				<div class="mdc-slider__track">
				  <div class="mdc-slider__track--inactive"></div>
				  <div class="mdc-slider__track--active">
					<div class="mdc-slider__track--active_fill"></div>
				  </div>
				</div>
				<div class="mdc-slider__thumb">
				  <div class="mdc-slider__thumb-knob"></div>
				</div>
			  </div>`)),
			),
		),
		mdc.JSInit(),
	)
}
