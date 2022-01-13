package mdc

type ButtonVariant int

const (
	ButtonRaised ButtonVariant = iota
	ButtonOutline
	ButtonText
)

func (bv ButtonVariant) ClassName() (class string) {
	switch bv {
	case ButtonRaised:
		class = "mdc-button--raised"

	case ButtonOutline:
		class = "mdc-button--outline"

	default:
		panic("unknown button variant")
	}
	return class
}
