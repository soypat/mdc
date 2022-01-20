package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/examples/jlog"
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

	hideSlider bool
}

func (b *Body) Render() vecty.ComponentOrHTML {
	dt := &mdc.DataTable{
		Columns: tabData,
		Rows:    tabRows,
	}
	return elem.Body(
		elem.Main(
			dt,
			elem.Div(
				&mdc.Button{
					Label: vecty.Text("Toggle Slider"),
					Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
						b.hideSlider = !b.hideSlider
						globalListener()
					})},
				},
				vecty.If(b.hideSlider, vecty.Text("slider hidden")),
				vecty.If(!b.hideSlider,
					elem.Div(
						&mdc.Slider{
							Name:    "slider-1",
							Max:     100,
							Variant: mdc.VariantSliderDiscrete,
						}),
				),
			),
		),
	)
}
