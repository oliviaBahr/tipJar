package models

import "tipJar/ui/styles"

type BaseComponent struct {
	Styler *styles.Styler
}

func NewBaseComponent() BaseComponent {
	return BaseComponent{
		Styler: styles.DefaultStyler,
	}
}
