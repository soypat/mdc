package mdc

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/soypat/mdc/examples/jlog"
)

type Button struct {
	vecty.Core // Do not modify.

	Label *vecty.HTML `vecty:"prop"`
	// TODO(soypat) Not implemented yet.
	// Size      Size                   `vecty:"prop"`
	Applyer    vecty.Applyer          `vecty:"prop"` // custom applier
	Style      ButtonStyle            `vecty:"prop"`
	Disabled   bool                   `vecty:"prop"`
	Listeners  []*vecty.EventListener `vecty:"prop"`
	Icon       IconType               `vecty:"prop"`
	ActionItem bool                   `vecty:"prop"`
}

func (b *Button) Render() vecty.ComponentOrHTML {
	jlog.Trace("Button.Render")
	hasIcon := b.Icon != ""
	markups := []vecty.Applyer{
		prop.Disabled(b.Disabled),
	}
	if b.Applyer == nil {
		// default applyer for loose buttons
		markups = append(markups, vecty.ClassMap{
			"mdc-button":               true,
			b.Style.ClassName():        true,
			"mdc-button--icon-leading": hasIcon,
		})
	} else {
		// custom applyer
		markups = append(markups, b.Applyer)
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
			vecty.Markup(
				vecty.Class("mdc-button__label"),
				// Accesibility for screen readers.
				vecty.MarkupIf(b.Label == nil, vecty.Property("aria-label", b.Icon.Name())),
			),
			vecty.If(b.Label != nil, b.Label),
		),
	)
}

func (b *Button) SetEventListeners(events ...*vecty.EventListener) *Button {
	b.Listeners = events
	return b
}

type Typography struct {
	vecty.Core // Do not modify.

	Applyer vecty.Applyer   `vecty:"prop"`
	Root    *vecty.HTML     `vecty:"prop"`
	Style   TypographyStyle `vecty:"prop"`
}

func (t *Typography) Render() vecty.ComponentOrHTML {
	jlog.Trace("Typography.Render")
	if t.Root == nil {
		panic("Root not set. did you forget to set it?")
	}
	element := t.Style.Element()
	return element(
		vecty.Markup(
			vecty.Class(t.Style.ClassName()),
			vecty.MarkupIf(t.Applyer != nil, t.Applyer),
		),
		t.Root,
	)
}

// Navbar represents a App bar in the top position.
type Navbar struct {
	vecty.Core

	Variant       AppBarVariant `vecty:"prop"`
	SectionStart  vecty.List    `vecty:"prop"`
	SectionCenter vecty.List    `vecty:"prop"`
	SectionEnd    vecty.List    `vecty:"prop"`
}

func (tb *Navbar) Render() vecty.ComponentOrHTML {
	jlog.Trace("TopBar.Render")
	for _, e := range tb.SectionStart {
		tb.apply(e)
	}
	for _, e := range tb.SectionCenter {
		tb.apply(e)
	}
	for _, e := range tb.SectionEnd {
		tb.apply(e)
	}
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

func (tb *Navbar) apply(sectionItem vecty.ComponentOrHTML) {
	switch e := sectionItem.(type) {
	case *Button:
		hasIcon := e.Icon != ""
		onlyIcon := hasIcon && e.Label == nil
		if onlyIcon {
			e.Applyer = vecty.ClassMap{
				"mdc-top-app-bar__" + e.Icon.Name() + "-icon": onlyIcon,
				"material-icons":               onlyIcon,
				"mdc-icon-button":              onlyIcon,
				"mdc-top-app-bar__action-item": e.ActionItem,
			}
		} else {
			// TODO(soypat): haven't seen what buttons should look like in
			e.Applyer = vecty.ClassMap{
				"mdc-button":                   true,
				e.Style.ClassName():            true,
				"mdc-top-app-bar__action-item": true,
			}
		}

	case *Typography:
		if e.Style.IsHeadline() {
			e.Applyer = vecty.Class("mdc-top-app-bar__title")
		}
	}
}

type icon struct {
	vecty.Core
	Kind    IconType `vecty:"prop"`
	Subtype string   `vecty:"prop"`
}

func (c *icon) Render() vecty.ComponentOrHTML {
	jlog.Trace("icon.Render")
	classes := vecty.ClassMap{
		"material-icons":              true,
		"mdc-" + c.Subtype + "__icon": c.Subtype != "",
	}
	return vecty.Tag("i",
		vecty.Markup(
			classes,
			vecty.Property("aria-hidden", true),
		),
		vecty.Text(c.Kind.Name()),
	)
}

func newButtonIcon(kind IconType) *icon {
	return &icon{
		Subtype: "button",
		Kind:    kind,
	}
}

// Leftbar is similar to the navbar but comes at you from a side
// of the application.
type Leftbar struct {
	vecty.Core
}

func (c *Leftbar) Render() vecty.ComponentOrHTML {
	list := &List{}
	return vecty.Tag("aside",
		vecty.Markup(vecty.Class("mdc-drawer")),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-drawer__content")),
			elem.Navigation(
				// TODO add list
				list,
			),
		),
	)
}

type List struct {
	vecty.Core
}

func (l *List) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList()
}
