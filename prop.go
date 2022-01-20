package mdc

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ButtonStyle int

const (
	defaultButtonStyle ButtonStyle = iota
	ButtonRaised
	ButtonOutline
	ButtonText
)

func (bs ButtonStyle) ClassName() (class string) {
	switch bs {
	case ButtonText:
		class = "mdc-button"
	case ButtonRaised, defaultButtonStyle:
		class = "mdc-button--raised"
	case ButtonOutline:
		class = "mdc-button--outline"
	default:
		panic("mdc: unknown button variant")
	}
	return class
}

type LeftbarVariant int

const (
	defaultLeftbarVariant LeftbarVariant = iota
	VariantStandardLeftbar
	VariantDismissableLeftbar
	VariantModalLeftbar
)

func (lb LeftbarVariant) IsDismissable() bool {
	return lb == VariantDismissableLeftbar || lb == VariantModalLeftbar
}

type Size int

const (
	defaultSize Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
)

func (s Size) Name() (str string) {
	switch s {
	case defaultSize, SizeSmall:
		str = "small"
	case SizeMedium:
		str = "medium"
	case SizeLarge:
		str = "large"
	default:
		panic("mdc: unknown Size property")
	}
	return str
}

type TypographyStyle int

const (
	defaultTypography TypographyStyle = iota
	Headline1
	Headline2
	Headline3
	Headline4
	Headline5
	Headline6
	Subtitle1
	Subtitle2
	Body1
	Body2
	Caption
)

func (ts TypographyStyle) ClassName() (class string) {
	class = "mdc-typography--"
	switch ts {
	case defaultTypography:
		class += "body1"

	case Headline1, Headline2, Headline3, Headline4, Headline5, Headline6:
		class += "headline" + strconv.Itoa(1+int(ts-Headline1))

	case Subtitle1, Subtitle2:
		class += "subtitle" + strconv.Itoa(1+int(ts-Subtitle1))

	case Body1, Body2:
		class += "body" + strconv.Itoa(1+int(ts-Body1))

	case Caption:
		class += "caption"

	default:
		panic("mdc: unknown typography variant")
	}
	return class
}

func (ts TypographyStyle) Element() (element func(markup ...vecty.MarkupOrChild) *vecty.HTML) {
	switch ts {
	case Headline1:
		element = elem.Heading1
	case Headline2:
		element = elem.Heading2
	case Headline3:
		element = elem.Heading3
	case Headline4:
		element = elem.Heading4
	case Headline5:
		element = elem.Heading5
	case Headline6:
		element = elem.Heading6
	case Caption:
		element = elem.Caption
	case Subtitle1, Subtitle2:
		element = elem.Subscript
	case Body1, Body2, defaultTypography:
		element = elem.Paragraph
	default:
		panic("mdc: unknown typography variant")
	}
	return element
}

func (ts TypographyStyle) IsHeadline() bool {
	return ts <= Headline6 && ts >= Headline1
}

type AppBarVariant int

const (
	variantTopBarDefault AppBarVariant = iota
	VariantTopBarShort
	VariantTopBarShortCollapsed
	VariantTopBarFixed
	VariantTopBarProminent
	VariantTopBarDense
)

const appBarID = "app-bar"

func (bv AppBarVariant) ClassName() (class string) {
	class = "mdc-top-app-bar"
	switch bv {
	case variantTopBarDefault:
		// No modifier to class.
	case VariantTopBarShort:
		class += "--short"
	case VariantTopBarShortCollapsed:
		class += "--short-collapsed"
	case VariantTopBarFixed:
		class += "--fixed"
	case VariantTopBarProminent:
		class += "--prominent"
	case VariantTopBarDense:
		class += "--dense"
	default:
		panic("mdc: unknown top bar variant")
	}
	return class
}

func (bv AppBarVariant) AdjustClassName() (class string) {
	return bv.ClassName() + "-adjust"
}

type ListElem int

const (
	defaultList ListElem = iota
	ElementUnorderedList
	ElementOrderedList
	ElementNavigationList
	ElementDivList
)

func (le ListElem) Element() (element func(markup ...vecty.MarkupOrChild) *vecty.HTML) {
	switch le {
	case ElementUnorderedList, defaultList:
		element = elem.UnorderedList
	case ElementDivList:
		element = elem.Div
	case ElementNavigationList:
		element = elem.Navigation
	case ElementOrderedList:
		element = elem.OrderedList
	default:
		panic("unknown ListElem")
	}
	return element
}

type ListItemElem int

const (
	defaultListItem ListItemElem = iota
	ElementSpanListItem
	ElementAnchorListItem
)

func (le ListItemElem) Element() (element func(markup ...vecty.MarkupOrChild) *vecty.HTML) {
	switch le {
	case ElementSpanListItem, defaultListItem:
		element = elem.Span
	case ElementAnchorListItem:
		element = elem.Anchor
	default:
		panic("unknown ListItemElem")
	}
	return element
}

type SliderVariant int

const (
	defaultSliderVariant SliderVariant = iota
	VariantSliderContinuous
	// Also continuous.
	VariantSliderRange
	VariantSliderDiscrete
)

func (sv SliderVariant) ClassName() (class string) {
	switch sv {
	case VariantSliderContinuous, defaultSliderVariant:
		// no specific class
	case VariantSliderRange:
		class = "mdc-slider--range"
	case VariantSliderDiscrete:
		class = "mdc-slider--discrete"
	default:
		panic("unknown SliderVariant")
	}
	return class
}

type DialogVariant int

const (
	defaultDialogVariant DialogVariant = iota
	// Full-screen dialogs group a series of tasks, such as creating a calendar entry with the event title, date, location, and time.
	VariantFullScreen
)

func (sv DialogVariant) ClassName() (class string) {
	switch sv {
	case VariantFullScreen, defaultDialogVariant:
		class = "mdc-dialog--fullscreen"
	default:
		panic("unknown DialogVariant")
	}
	return class
}
