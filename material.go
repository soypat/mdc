package mdc

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/soypat/mdc/examples/jlog"
)

type Button struct {
	vecty.Core // Do not modify.

	Label     *vecty.HTML            `vecty:"prop"`
	Style     ButtonStyle            `vecty:"prop"`
	Disabled  bool                   `vecty:"prop"`
	Listeners []*vecty.EventListener `vecty:"prop"`
	Icon      string                 `vecty:"prop"`
}

func (b *Button) Render() vecty.ComponentOrHTML {
	jlog.Trace("Button.Render")
	hasIcon := b.Icon != ""
	classes := vecty.ClassMap{
		"mdc-button":               true,
		b.Style.ClassName():        true,
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

type Typography struct {
	vecty.Core // Do not modify.

	Root  *vecty.HTML
	Style TypographyStyle
}

func (t *Typography) Render() vecty.ComponentOrHTML {
	jlog.Trace("Typography.Render")
	if t.Root == nil {
		panic("Root not set. did you forget to set it?")
	}
	element := t.Style.Element()
	return element(
		vecty.Markup(vecty.Class(t.Style.ClassName())),
		t.Root,
	)
}

// TopBar represents a App bar in the top position.
type TopBar struct {
	vecty.Core

	Variant       AppBarVariant
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
}

func (tb *TopBar) Render() vecty.ComponentOrHTML {
	jlog.Trace("TopBar.Render")
	return elem.Header(vecty.Markup(vecty.Class("mdc-top-app-bar", tb.Variant.ClassName())),

		elem.Div(vecty.Markup(vecty.Class("mdc-top-app-bar__row")),
			//Div contents
			vecty.If(len(tb.SectionStart) != 0,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
						vecty.Class("mdc-top-app-bar__section--align-start"),
					),
					tb.SectionStart,
				),
			),
			vecty.If(len(tb.SectionCenter) != 0,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
						vecty.Class("mdc-top-app-bar__section--align"),
					),
					tb.SectionCenter,
				),
			),
			vecty.If(len(tb.SectionEnd) != 0,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
						vecty.Class("mdc-top-app-bar__section--align-end"),
					),
					tb.SectionEnd,
				),
			),
		),
	)
}

// AdjustmentClass returns the class name for <main> element to adjust
// view.
func (tb *TopBar) AdjustmentClass() string {
	return tb.Variant.ClassName() + "-adjust"
}
