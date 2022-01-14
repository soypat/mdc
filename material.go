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
	Icon      IconType               `vecty:"prop"`
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
		vecty.If(
			hasIcon, newButtonIcon(b.Icon),
		),
		elem.Span(
			vecty.Markup(vecty.Class("mdc-button__label")),
			b.Label,
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

// Navbar represents a App bar in the top position.
type Navbar struct {
	vecty.Core

	Variant       AppBarVariant
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
}

func (tb *Navbar) Render() vecty.ComponentOrHTML {
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
func (tb *Navbar) AdjustmentClass() string {
	return tb.Variant.ClassName() + "-adjust"
}

type icon struct {
	vecty.Core
	kind    IconType
	subtype string
}

func (c *icon) Render() vecty.ComponentOrHTML {
	jlog.Trace("icon.Render")
	classes := vecty.ClassMap{
		"material-icons":              true,
		"mdc-" + c.subtype + "__icon": c.subtype != "",
	}
	return vecty.Tag("i",
		vecty.Markup(
			classes,
			vecty.Property("aria-hidden", true),
		),
		vecty.Text(c.kind.Name()),
	)
}

func newButtonIcon(kind IconType) *icon {
	return &icon{
		subtype: "button",
		kind:    kind,
	}
}
