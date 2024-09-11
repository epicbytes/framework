package admin

import (
	"github.com/a-h/templ"
	"github.com/samber/lo"
)

/* Checkbox partial templ file for input.templ */

type CheckboxModifier string
type CheckboxResponsive string

const (
	// CheckboxModifier_Primary adds `primary` to checkbox
	CheckboxModifier_Primary CheckboxModifier = "checkbox-primary"
	// CheckboxModifier_Secondary adds `secondary` to checkbox
	CheckboxModifier_Secondary CheckboxModifier = "checkbox-secondary"
	// CheckboxModifier_Accent adds `accent` to checkbox
	CheckboxModifier_Accent CheckboxModifier = "checkbox-accent"
	// CheckboxModifier_Success adds `success` to checkbox
	CheckboxModifier_Success CheckboxModifier = "checkbox-success"
	// CheckboxModifier_Warning adds `warning` to checkbox
	CheckboxModifier_Warning CheckboxModifier = "checkbox-warning"
	// CheckboxModifier_Info adds `info` to checkbox
	CheckboxModifier_Info CheckboxModifier = "checkbox-info"
	// CheckboxModifier_Error adds `error` to checkbox
	CheckboxModifier_Error CheckboxModifier = "checkbox-error"
)

const (
	// CheckboxResponsive_Lg For Large size for checkbox
	CheckboxResponsive_Lg CheckboxResponsive = "checkbox-lg"
	// CheckboxResponsive_Md For Medium (default) size for checkbox
	CheckboxResponsive_Md CheckboxResponsive = "checkbox-md"
	// CheckboxResponsive_Sm For Small size for checkbox
	CheckboxResponsive_Sm CheckboxResponsive = "checkbox-sm"
	// CheckboxResponsive_Xs For Extra small size for checkbox
	CheckboxResponsive_Xs CheckboxResponsive = "checkbox-xs"
)

func WithCheckboxModifiers(checkboxModifier ...CheckboxModifier) InputOption {
	return func(b *input) {
		b.checkboxModifier = lo.Uniq(append(b.checkboxModifier, checkboxModifier...))
	}
}

func WithCheckboxResponsive(checkboxResponsive ...CheckboxResponsive) InputOption {
	return func(b *input) {
		b.checkboxResponsive = lo.Uniq(append(b.checkboxResponsive, checkboxResponsive...))
	}
}

/*
NewCheckboxSm creates checkbox smaller with some props
Inputs allow the user to input values.
*/
func NewCheckboxSm(name string, options ...InputOption) templ.Component {
	return NewInput(name, append(options, WithInputValue("true"), WithInputType(InputType_Checkbox), WithCheckboxResponsive(CheckboxResponsive_Sm))...)
}
