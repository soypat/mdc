package mdc

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type Button struct {
	vecty.Core

	Label     *vecty.HTML            `vecty:"prop"`
	Variant   ButtonVariant          `vecty:"prop"`
	Disabled  bool                   `vecty:"prop"`
	Listeners []*vecty.EventListener `vecty:"prop"`
	Icon      string                 `vecty:"prop"`
}

func (b *Button) Render() vecty.ComponentOrHTML {
	hasIcon := b.Icon != ""
	classes := vecty.ClassMap{
		"mdc-button":               true,
		b.Variant.ClassName():      true,
		"mdc-button--icon-leading": hasIcon,
	}
	markups := []vecty.Applyer{
		prop.Disabled(b.Disabled),
		classes,
	}
	for i := range b.Listeners {
		markups = append(markups, b.Listeners[i])
	}
	return elem.Button(
		vecty.Markup(markups...),
		elem.Span(
			vecty.Markup(vecty.Class("mdc-button__label")),
			b.Label,
		),
		vecty.If(
			hasIcon, vecty.Tag("i"),
		),
	)
}

func (b *Button) SetEventListeners(events ...*vecty.EventListener) *Button {
	b.Listeners = events
	return b
}
