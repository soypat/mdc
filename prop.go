package mdc

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ButtonStyle int

const (
	ButtonRaised ButtonStyle = iota
	ButtonOutline
	ButtonText
)

func (bs ButtonStyle) ClassName() (class string) {
	switch bs {
	case ButtonRaised:
		class = "mdc-button--raised"
	case ButtonOutline:
		class = "mdc-button--outline"
	default:
		panic("unknown button variant")
	}
	return class
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
		panic("unknown typography variant")
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
	case Body1, Body2:
		element = elem.Body
	default:
		panic("unknown typography variant")
	}
	return element
}

type AppBarVariant int

const (
	VariantTopBarShort AppBarVariant = iota
	VariantTopBarShortCollapsed
	VariantTopBarFixed
	VariantTopBarProminent
	VariantTopBarDense
)

func (bv AppBarVariant) ClassName() (class string) {
	class = "mdc-top-app-bar--"
	switch bv {
	case VariantTopBarShort:
		class += "short"
	case VariantTopBarShortCollapsed:
		class += "short-collapsed"
	case VariantTopBarFixed:
		class += "fixed"
	case VariantTopBarProminent:
		class += "prominent"
	case VariantTopBarDense:
		class += "dense"
	default:
		panic("unknown top bar variant")
	}
	return class
}

type IconType string

// TODO(soypat) add more MDC icons.
const (
	IconBookmark IconType = "bookmark"
)

func (c IconType) Name() string {
	return string(c)
}
