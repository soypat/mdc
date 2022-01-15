package mdc

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
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
	hasIcon := b.Icon.IsValid()
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
		hasIcon := e.Icon.IsValid()
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

	Title       *vecty.HTML `vecty:"prop"`
	Subtitle    *vecty.HTML `vecty:"prop"`
	List        *List       `vecty:"prop"`
	Dismissible bool        `vecty:"prop"`
	Closed      bool        `vecty:"prop"`
}

func (c *Leftbar) Render() vecty.ComponentOrHTML {
	c.List.ListElem = ElementNavigation
	hasHeader := c.Title != nil || c.Subtitle != nil
	return vecty.Tag("aside",
		vecty.Markup(
			vecty.Class("mdc-drawer"),
			vecty.MarkupIf(c.Dismissible, vecty.Class("mdc-drawer--dismissible")),
			vecty.MarkupIf(!c.Closed, vecty.Class("mdc-drawer--open")),
		),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-drawer__content")),
			vecty.If(hasHeader,
				elem.Div(
					vecty.Markup(vecty.Class("mdc-drawer__header")),
					vecty.If(c.Title != nil, elem.Heading3(vecty.Markup(vecty.Class("mdc-drawer__title")),
						c.Title,
					)),
					vecty.If(c.Subtitle != nil, elem.Heading3(vecty.Markup(vecty.Class("mdc-drawer__subtitle")),
						c.Subtitle,
					)),
				),
			),
			c.List,
		),
	)
}

type List struct {
	vecty.Core

	List          vecty.List                    `vecty:"prop"`
	ListElem      ListElem                      `vecty:"prop"`
	ClickListener func(idx int, e *vecty.Event) `vecty:"prop"`
}

func (l *List) Render() vecty.ComponentOrHTML {
	element := l.ListElem.Element()
	if l.ClickListener != nil {
		for i, v := range l.List {
			i := i // escape loop variable
			switch e := v.(type) {
			case *ListItem:
				e.Listeners = append(e.Listeners, event.Click(func(e *vecty.Event) {
					l.ClickListener(i, e)
				}))
			}
		}
	}

	return element(vecty.Markup(vecty.Class("mdc-list")),
		l.List,
	)
}

type ListItem struct {
	vecty.Core

	Label        *vecty.HTML            `vecty:"prop"`
	Icon         IconType               `vecty:"prop"`
	ListItemElem ListItemElem           `vecty:"prop"`
	Active       bool                   `vecty:"prop"`
	Listeners    []*vecty.EventListener `vecty:"prop"`
}

func (l *ListItem) Render() vecty.ComponentOrHTML {
	hasIcon := l.Icon.IsValid()
	element := l.ListItemElem.Element()
	listeners := make([]vecty.Applyer, len(l.Listeners))
	for i := range l.Listeners {
		listeners[i] = l.Listeners[i]
	}
	return element(
		vecty.Markup(
			vecty.Class("mdc-list-item"),
			vecty.MarkupIf(l.Active, vecty.Class("mdc-list-item--activated")),
			vecty.MarkupIf(len(listeners) > 0, listeners...),
		),

		vecty.If(hasIcon,
			vecty.Tag("i", vecty.Markup(vecty.Class("material-icons", "mdc-list-item__graphic")),
				vecty.Text(l.Icon.Name()),
			),
		),

		vecty.If(l.Label != nil,
			elem.Span(
				vecty.Markup(vecty.Class("mdc-list-item__text")),
				l.Label,
			),
		),
	)
}

var badFormID = "form inputs require unique, non empty IDs"

// Checkboxes should always rendered as indeterminate on startup.
type Checkbox struct {
	vecty.Core
	ID       string      `vecty:"prop"`
	Label    *vecty.HTML `vecty:"prop"`
	Disabled bool        `vecty:"prop"`
	NoRipple bool        `vecty:"prop"`
}

func (cb *Checkbox) Render() vecty.ComponentOrHTML {
	if cb.ID == "" {
		panic(badFormID)
	}
	return elem.Div(vecty.Markup(vecty.Class("mdc-form-field")),
		elem.Div(vecty.Markup(
			vecty.Class("mdc-checkbox"),
			vecty.MarkupIf(cb.Disabled, vecty.Class("mdc-checkbox--disabled")),
		),
			elem.Input(vecty.Markup(
				prop.Type(prop.TypeCheckbox),
				prop.ID(cb.ID),
				vecty.Class("mdc-checkbox__native-control"),
				// TODO(soypat): How to implement data determination.
				vecty.MarkupIf(false, vecty.Attribute("data-indeterminate", "true")),
			)),
			elem.Div(
				vecty.Markup(vecty.Class("mdc-checkbox__background"), vecty.UnsafeHTML(svgCheckbox)),
			),
			elem.Div(vecty.Markup(vecty.MarkupIf(!cb.NoRipple), vecty.Class("mdc-checkbox__ripple"))),
		),
		elem.Label(vecty.Markup(prop.For(cb.ID)), cb.Label),
	)
}

type DataTable struct {
	vecty.Core
	Columns []Series `vecty:"prop"`
	Rows    int      `vecty:"prop"`
}

func (dt *DataTable) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("mdc-data-table")),
		elem.Div(vecty.Markup(vecty.Class("mdc-data-table__table-container")),
			elem.Table(vecty.Markup(vecty.Class("mdc-data-table__table")),
				elem.TableHead(elem.TableRow(
					vecty.Markup(vecty.Class("mdc-data-table__header-row")),
					dt.heads()),
				),
			),
			vecty.If(dt.Rows > 0, elem.TableBody(vecty.Markup(vecty.Class("mdc-data-table__content")),
				dt.rows(),
			)),
		),
	)
}

func (dt *DataTable) heads() vecty.MarkupOrChild {
	var heads vecty.List
	for i := 0; i < len(dt.Columns); i++ {
		k := dt.Columns[i].Kind()
		switch k {
		case DataString, DataNumeric:
			heads = append(heads, elem.TableHeader(vecty.Markup(
				vecty.Class("mdc-data-table__header-cell"),
				vecty.Attribute("role", "columnheader"),
				vecty.Attribute("scope", "col"),
				vecty.MarkupIf(k == DataNumeric, vecty.Class(k.ClassName())),
			),
				vecty.Text(dt.Columns[i].Head()),
			))
		}
	}
	return heads
}

func (dt *DataTable) rows() vecty.MarkupOrChild {
	rows := make(vecty.List, dt.Rows)
	thMarkup := []vecty.Applyer{vecty.Class("mdc-data-table__cell"), vecty.Attribute("scope", "row")}
	tdMarkup := []vecty.Applyer{vecty.Class("mdc-data-table__cell")}
	for i := 0; i < dt.Rows; i++ {
		row := make(vecty.List, len(dt.Columns))
		for j := range dt.Columns {
			markup := tdMarkup
			if j == 0 { // append as header
				markup = thMarkup
			}
			row = append(row, elem.TableHeader(
				vecty.Markup(markup...),
				dt.Columns[j].AtRow(i),
			))
		}
		rows = append(rows, elem.TableRow(vecty.Markup(vecty.Class("mdc-data-table__row")),
			row,
		))
	}
	return rows
}
