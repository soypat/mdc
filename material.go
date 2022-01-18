package mdc

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/soypat/mdc/examples/jlog"
	"github.com/soypat/mdc/icons"
)

// Button implements the button Material Design Component
// https://material.io/components/buttons
//
// Take care when adding buttons in <form> elements as they
// may act as Submit actions if not avoided correctly.
type Button struct {
	vecty.Core // Do not modify.

	Label *vecty.HTML `vecty:"prop"`
	// TODO(soypat) Not implemented yet.
	// Size      Size                   `vecty:"prop"`
	Applyer    vecty.Applyer          `vecty:"prop"` // custom applier
	Style      ButtonStyle            `vecty:"prop"`
	Disabled   bool                   `vecty:"prop"`
	Listeners  []*vecty.EventListener `vecty:"prop"`
	Icon       icons.Icon             `vecty:"prop"`
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

// Typography implements Material Design's Typography element
// https://material.io/design/typography/the-type-system.html.
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
// https://material.io/components/app-bars-top.
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
	return elem.Header(vecty.Markup(vecty.Class("app-bar", "mdc-top-app-bar", tb.Variant.ClassName()), prop.ID(appBarID)),

		elem.Div(vecty.Markup(
			vecty.Class("mdc-top-app-bar__row"),
		),

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
	Kind    icons.Icon `vecty:"prop"`
	Subtype string     `vecty:"prop"`
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

func newButtonIcon(kind icons.Icon) *icon {
	return &icon{
		Subtype: "button",
		Kind:    kind,
	}
}

// Leftbar implements the navigation drawer Material Design component
// https://material.io/components/navigation-drawer.
type Leftbar struct {
	vecty.Core

	Title       *vecty.HTML    `vecty:"prop"`
	Subtitle    *vecty.HTML    `vecty:"prop"`
	List        *List          `vecty:"prop"`
	Variant     LeftbarVariant `vecty:"prop"`
	StartClosed bool           `vecty:"prop"`
	Applyer     vecty.Applyer  `vecty:"prop"`
	NoJS        bool           `vecty:"prop"`
}

func (lb *Leftbar) id() string { return ".mdc-drawer" }

func (lb *Leftbar) Render() vecty.ComponentOrHTML {
	jlog.Trace("Leftbar.Render")
	dismissable := lb.Variant.IsDismissable()
	if lb.List.ListElem == defaultList {
		jlog.Trace("Leftbar Listelem autoset to nav")
		lb.List.ListElem = ElementNavigationList
	}
	hasHeader := lb.Title != nil || lb.Subtitle != nil
	return vecty.Tag("aside",
		vecty.Markup(
			vecty.Class("mdc-drawer"),
			vecty.MarkupIf(lb.Applyer != nil, lb.Applyer),
			vecty.MarkupIf(dismissable, vecty.Class("mdc-drawer--dismissible")),
			vecty.MarkupIf(!lb.StartClosed, vecty.Class("mdc-drawer--open")),
		),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-drawer__content")),
			vecty.If(hasHeader,
				elem.Div(
					vecty.Markup(vecty.Class("mdc-drawer__header")),
					vecty.If(lb.Title != nil, elem.Heading3(vecty.Markup(vecty.Class("mdc-drawer__title")),
						lb.Title,
					)),
					vecty.If(lb.Subtitle != nil, elem.Heading3(vecty.Markup(vecty.Class("mdc-drawer__subtitle")),
						lb.Subtitle,
					)),
				),
			),
			lb.List,
		),
	)
}

func (lb *Leftbar) Mount() {
	jlog.Trace("Leftbar.Mount")
	if lb.Variant.IsDismissable() {
		handler := nsDrawer.newFromQuery("MDCDrawer", lb.id())
		err := globalHandlers.registerID(lb.id(), handler)
		jlog.Trace("Leftbar.Mount success", err)
	}
}

func (lb *Leftbar) SkipRender(c vecty.Component) bool {
	skip := !Handler(lb).IsUndefined()
	jlog.Trace("Leftbar.SkipRender()=>", skip)
	return skip
}

func (lb *Leftbar) Unmount() {
	jlog.Trace("Leftbar.Unmount")
	if lb.Variant.IsDismissable() {
		destroyHandler(lb)
		jlog.Trace("Unmount.destroy success")
	}
}

// Dismiss Only supposed to be used with javascript.
func (lb *Leftbar) Dismiss(closed bool) {
	Handler(lb).Set("open", !closed)
}

func (lb *Leftbar) IsDismissed() (closed bool) {
	return !Handler(lb).Get("open").Bool()
}

// List implements the list Material Design component
// https://material.io/components/lists.
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

// ListItem is a component for use with the List component.
type ListItem struct {
	vecty.Core

	Label        *vecty.HTML            `vecty:"prop"`
	Href         string                 `vecty:"prop"`
	Icon         icons.Icon             `vecty:"prop"`
	ListItemElem ListItemElem           `vecty:"prop"`
	Active       bool                   `vecty:"prop"`
	Listeners    []*vecty.EventListener `vecty:"prop"`
}

func (l *ListItem) Render() vecty.ComponentOrHTML {
	hasIcon := l.Icon.IsValid()
	isAnchor := l.ListItemElem == ElementAnchorListItem && l.Href != ""
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
			vecty.MarkupIf(isAnchor, prop.Href(l.Href)),
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

var badFormID = "form inputs require unique, non empty IDs or names"

// Checkbox implements the checkbox Material Design component
// https://material.io/components/checkboxes.
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

// DataTable implements the data table Material Design component
// https://material.io/components/data-tables
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
				vecty.If(dt.Rows > 0, elem.TableBody(vecty.Markup(vecty.Class("mdc-data-table__content")),
					dt.rows(),
				)),
			),
		),
	)
}

func (dt *DataTable) heads() vecty.MarkupOrChild {
	var heads vecty.List
	for i := 0; i < len(dt.Columns); i++ {
		k := dt.Columns[i].Kind()
		cc := k.CellClassName()
		cellClass := vecty.MarkupIf(cc != "", vecty.Class(cc))
		switch k {
		case DataString, DataNumeric:
			heads = append(heads, elem.TableHeader(vecty.Markup(
				vecty.Class("mdc-data-table__header-cell"),
				vecty.Attribute("role", "columnheader"),
				vecty.Attribute("scope", "col"),
				cellClass,
			),
				vecty.Text(dt.Columns[i].Head()),
			))
		case DataCheckbox:
			panic("checkbox DataKind for DataTable not implemented")
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

// Slider implements the slider Material Design component
// https://material.io/components/sliders/web#sliders.
// TODO(soypat): What is going on with this component?
// The examples in the online documentation do not yield a working slider?
type Slider struct {
	vecty.Core
	Name    string        `vecty:"prop"`
	Min     int           `vecty:"prop"`
	Max     int           `vecty:"prop"`
	Value   int           `vecty:"prop"`
	Step    int           `vecty:"prop"`
	Variant SliderVariant `vecty:"prop"`
}

func (s *Slider) id() string { return s.Name }

func (s *Slider) Mount() {
	jlog.Trace("Slider.Mount")
	handle := nsSlider.newFromId("MDCSlider", s.Name)
	err := globalHandlers.registerID(s.id(), handle)
	if err != nil {
		jlog.Debug("Slider.Mount register error", err)
	}
}

func (s *Slider) SkipRender(prev vecty.Component) bool {
	skip := !Handler(s).IsUndefined()
	jlog.Trace("Slider.SkipRender() =>", skip)
	return skip
}

func (s *Slider) Unmount() {
	jlog.Trace("Slider.Unmount")
	destroyHandler(s)
}

func (s *Slider) Render() vecty.ComponentOrHTML {
	jlog.Trace("slider.Render")
	if s.Name == "" {
		panic(badFormID)
	}
	if s.Value < s.Min || s.Value > s.Max {
		s.Value = (s.Min + s.Max) / 2 // Sane default: start at midpoint.
	}
	if s.Variant == VariantSliderDiscrete && s.Step == 0 {
		s.Step = (s.Max - s.Min) / 20 // 20 step increments for discrete sliders.
	}
	variantClass := s.Variant.ClassName()
	return elem.Div(vecty.Markup(
		prop.ID(s.Name),
		vecty.ClassMap{
			"mdc-slider": true,
			variantClass: variantClass != "",
		},
	),
		s.inputs(),
		s.display(),
	)
}

func (s *Slider) inputs() (inputs vecty.MarkupOrChild) {
	defaultMarkup := []vecty.Applyer{
		vecty.Class("mdc-slider__input"),
		vecty.Property("min", s.Min),
		vecty.Property("max", s.Max),
		vecty.Attribute("value", s.Value),
		prop.Type(prop.TypeRange),
		prop.Name(s.Name),
		vecty.MarkupIf(s.Step > 0, vecty.Property("step", s.Step)),
	}
	switch s.Variant {
	case VariantSliderContinuous, VariantSliderDiscrete, VariantSliderRange:
		// Single input
		inputs = elem.Input(
			vecty.Markup(defaultMarkup...),
		)
	default:
		panic("SliderVariant inputs not implemented")
	}
	return inputs
}

func (s *Slider) display() (disp vecty.MarkupOrChild) {
	switch s.Variant {
	case VariantSliderContinuous, VariantSliderDiscrete:
		// Single input
		disp = vecty.List{
			elem.Div(vecty.Markup(vecty.Class("mdc-slider__track")),
				elem.Div(vecty.Markup(vecty.Class("mdc-slider__track--inactive"))),
				elem.Div(vecty.Markup(vecty.Class("mdc-slider__track--active")),
					elem.Div(vecty.Markup(vecty.Class("mdc-slider__track--active_fill"))),
				),
			),
			//Thumb slide element
			elem.Div(vecty.Markup(vecty.Class("mdc-slider__thumb")),
				vecty.If(s.Variant == VariantSliderDiscrete,
					// Discrete Slider display bar.
					elem.Div(vecty.Markup(vecty.Class("mdc-slider__value-indicator-container"), vecty.Attribute("aria-hidden", true)),
						elem.Div(vecty.Markup(vecty.Class("mdc-slider__value-indicator")),
							elem.Span(vecty.Markup(vecty.Class("mdc-slider__value-indicator-text")),
								vecty.Text(strconv.Itoa(s.Value)),
							),
						),
					),
				),
				elem.Div(vecty.Markup(vecty.Class("mdc-slider__thumb-knob"))),
			),
		}
	default:
		panic("SliderVariant display not implemented")
	}
	return disp
}

// Tooltip implements the tooltip Material Design component
// https://material.io/components/tooltips/web#tooltips.
//
// From Material documentation: To ensure proper positioning of the tooltip,
// it's important that this tooltip element is an immediate child of the <body>,
// rather than nested underneath the anchor element or other elements.
type Tooltip struct {
	vecty.Core
	ID    string      `vecty:"prop"`
	Label *vecty.HTML `vecty:"prop"`
}

func (tt *Tooltip) id() string { return tt.ID }

func (tt *Tooltip) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Attribute("aria-describedby", tt.ID),
	).Apply(h)
}

func (tt *Tooltip) Render() vecty.ComponentOrHTML {
	if !mdcOK() {
		panic("MDC script required to use Tooltips")
	}
	return elem.Div(vecty.Markup(
		prop.ID(tt.ID),
		vecty.Class("mdc-tooltip"),
		vecty.Attribute("role", "tooltip"),
		vecty.Attribute("aria-hidden", true)),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-tooltip__surface", "mdc-tooltip__surface-animation")),
			tt.Label,
		),
	)
}

// Mount TODO(soypat): is tricky business. Javascript must be called on the HTML node in the DOM
// after it is rendered. Luckily tooltip should be static if not removed from the dom.
// This probably breaks if tooltip is removed by vecty. Must find JS methods for "destroying"
// the javascript handle to the tooltip for use with the vecty.Unmount interface.
//
// DO NOT CALL MOUNT METHOD YOURSELF. This method is called by vecty only.
func (tt *Tooltip) Mount() {
	handler := nsTooltip.newFromId("MDCTooltip", tt.ID)
	globalHandlers.registerID(tt.id(), handler)
}

func (tt *Tooltip) Unmount() {
	destroyHandler(tt)
}

// SPA implements the suggested combination of Appbar with
// a drawer component according to Material Design guidelines
// https://material.io/components/navigation-drawer/web
// (see Usage with top app bar)
type SPA struct {
	vecty.Core
	Navbar           *Navbar             `vecty:"prop"`
	Drawer           *Leftbar            `vecty:"prop"`
	Content          vecty.MarkupOrChild `vecty:"prop"`
	FullHeightDrawer bool                `vecty:"prop"`
}

func (spa *SPA) Render() vecty.ComponentOrHTML {
	style := elem.Style(vecty.Markup(vecty.UnsafeHTML(spaCSS)))
	spa.Drawer.List.ListElem = ElementDivList
	nbAdjust := VariantTopBarFixed.AdjustClassName()
	if spa.FullHeightDrawer {
		return elem.Div(
			style,
			spa.Drawer,
			elem.Div(vecty.Markup(vecty.Class("mdc-drawer-app-content")),
				spa.Navbar,
				elem.Main(vecty.Markup(
					vecty.Class("main-content"),
					prop.ID("main-content"),
				),
					elem.Div(vecty.Markup(vecty.Class(nbAdjust)),
						spa.Content,
					),
				),
			),
		)
	}
	// TODO(soypat): Fix this branch. Drawer is being
	// rendered as full height. Might have missed something subtle...
	spa.Drawer.Applyer = vecty.Class(nbAdjust)
	return elem.Div(
		style,
		spa.Navbar,
		spa.Drawer,
		elem.Div(vecty.Markup(vecty.Class("mdc-drawer-app-content", nbAdjust)),
			elem.Main(vecty.Markup(
				vecty.Class("main-content"),
				prop.ID("main-content"),
			),
				spa.Content,
			),
		),
	)
}
