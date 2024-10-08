// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"fmt"
	"github.com/epicbytes/framework/webui/shared"
	"github.com/samber/lo"
	"strings"
)

type InputType string
type InputModifier string
type InputResponsive string

const (
	InputType_Button        InputType = "button"
	InputType_Checkbox      InputType = "checkbox"
	InputType_Color         InputType = "color"
	InputType_Date          InputType = "date"
	InputType_DatetimeLocal InputType = "datetime-local"
	InputType_Email         InputType = "email"
	InputType_File          InputType = "file"
	InputType_Hidden        InputType = "hidden"
	InputType_Image         InputType = "image"
	InputType_Month         InputType = "month"
	InputType_Number        InputType = "number"
	InputType_Password      InputType = "password"
	InputType_Radio         InputType = "radio"
	InputType_Range         InputType = "range"
	InputType_Reset         InputType = "reset"
	InputType_Search        InputType = "search"
	InputType_Submit        InputType = "submit"
	InputType_Tel           InputType = "tel"
	InputType_Text          InputType = "text"
	InputType_Time          InputType = "time"
	InputType_Url           InputType = "url"
	InputType_Week          InputType = "week"
)

const (
	// InputModifier_Bordered adds border to input
	InputModifier_Bordered InputModifier = "input-bordered"
	// InputModifier_Ghost adds ghost style to input
	InputModifier_Ghost InputModifier = "input-ghost"
	// InputModifier_Primary adds `primary` color to input
	InputModifier_Primary InputModifier = "input-primary"
	// InputModifier_Secondary adds `secondary` color to input
	InputModifier_Secondary InputModifier = "input-secondary"
	// InputModifier_Accent adds `accent` color to input
	InputModifier_Accent InputModifier = "input-accent"
	// InputModifier_Info adds `info` color to input
	InputModifier_Info InputModifier = "input-info"
	// InputModifier_Success adds `success` color to input
	InputModifier_Success InputModifier = "input-success"
	// InputModifier_Warning adds `warning` color to input
	InputModifier_Warning InputModifier = "input-warning"
	// InputModifier_Error adds `error` color to input
	InputModifier_Error InputModifier = "input-error"
)

const (
	// InputResponsive_Lg For Large size for input
	InputResponsive_Lg InputResponsive = "input-lg"
	// InputResponsive_Md For Medium (default) size for input
	InputResponsive_Md InputResponsive = "input-md"
	// InputResponsive_Sm For Small size for input
	InputResponsive_Sm InputResponsive = "input-sm"
	// InputResponsive_Xs For Extra small size for input
	InputResponsive_Xs InputResponsive = "input-xs"
)

type input struct {
	label              string
	name               string
	placeholder        string
	value              string
	inputChecked       bool
	classes            templ.CSSClasses
	inputType          InputType
	inputModifier      []InputModifier
	inputResponsive    []InputResponsive
	checkboxModifier   []CheckboxModifier
	checkboxResponsive []CheckboxResponsive
	icon               shared.SpritesBase
	attrs              templ.Attributes
}

type InputOption func(*input)

func (b *input) buildClasses() templ.CSSClasses {
	var classes string

	classes = b.classes.String()

	if b.inputType == InputType_Checkbox || b.inputType == InputType_Radio {
		classes += " " + strings.Join(lo.Map(b.checkboxModifier, func(item CheckboxModifier, index int) string {
			return string(item)
		}), " ")

		classes += " " + strings.Join(lo.Map(b.checkboxResponsive, func(item CheckboxResponsive, index int) string {
			return string(item)
		}), " ")
	} else {
		classes += " " + strings.Join(lo.Map(b.inputModifier, func(item InputModifier, index int) string {
			return string(item)
		}), " ")

		classes += " " + strings.Join(lo.Map(b.inputResponsive, func(item InputResponsive, index int) string {
			return string(item)
		}), " ")
	}
	unifiedClasses := lo.Uniq(strings.Split(classes, " "))
	return templ.Classes(unifiedClasses)
}

func (b *input) getError(ctx context.Context) string {
	if values, ok := ctx.Value("errors").(map[string]string); ok {
		return lo.ValueOr(values, b.name, "")
	} else {
		return ""
	}
}

func WithInputValue(value any) InputOption {
	return func(b *input) {
		b.value = fmt.Sprintf("%v", value)
	}
}

func WithInputChecked(inputChecked bool) InputOption {
	return func(b *input) {
		b.inputChecked = inputChecked
	}
}

func WithInputType(inputType InputType) InputOption {
	return func(b *input) {
		b.inputType = inputType
	}
}

func WithInputLabel(label string) InputOption {
	return func(b *input) {
		b.label = label
	}
}

func WithInputPlaceholder(placeholder string) InputOption {
	return func(b *input) {
		b.placeholder = placeholder
	}
}

func WithInputModifiers(inputModifier ...InputModifier) InputOption {
	return func(b *input) {
		b.inputModifier = lo.Uniq(append(b.inputModifier, inputModifier...))
	}
}

func WithInputResponsive(inputResponsive ...InputResponsive) InputOption {
	return func(b *input) {
		b.inputResponsive = lo.Uniq(append(b.inputResponsive, inputResponsive...))
	}
}

func WithInputAttrs(attrs templ.Attributes) InputOption {
	return func(b *input) {
		b.attrs = attrs
	}
}

func (b *input) render() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if b.inputType == InputType_Hidden {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"hidden\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(b.value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 174, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(b.name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 174, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, b.attrs)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if b.inputType == InputType_Checkbox || b.inputType == InputType_Radio {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"form-control w-auto\"><label class=\"cursor-pointer label flex items-center justify-start\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 = []any{"checkbox", b.buildClasses(), templ.KV("checkbox-error", len(b.getError(ctx)) > 0)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"checkbox\" name=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(b.name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 178, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if b.inputChecked {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var4).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, b.attrs)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> <span class=\"pl-3 label-text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(b.label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 179, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(b.getError(ctx)) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"label\"><span class=\"label-text-alt text-error\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(b.getError(ctx))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 183, Col: 62}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"form-control w-full\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(b.label) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"label\"><span class=\"label-text\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(b.label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 191, Col: 39}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var10 = []any{"input", b.buildClasses(), templ.KV("input-error", len(b.getError(ctx)) > 0)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var10...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(string(b.inputType))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 194, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(b.value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 194, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(b.name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 194, Col: 70}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" placeholder=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(b.placeholder)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 194, Col: 100}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var10).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, b.attrs)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(b.getError(ctx)) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"label\"><span class=\"label-text-alt text-error\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var16 string
				templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(b.getError(ctx))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/input.templ`, Line: 197, Col: 62}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

/*
NewInput creates input with some props
*/
func NewInput(name string, options ...InputOption) templ.Component {
	inp := &input{
		name:          name,
		inputType:     InputType_Text,
		inputModifier: []InputModifier{InputModifier_Bordered},
	}

	for _, opt := range options {
		opt(inp)
	}

	return inp.render()
}

/*
NewInputSm creates input smaller with some props
Inputs allow the user to input values.
*/
func NewInputSm(name string, options ...InputOption) templ.Component {
	return NewInput(name, append(options, WithInputResponsive(InputResponsive_Sm))...)
}

func NewHiddenInput(name string, value any) templ.Component {
	return NewInput(name, WithInputType(InputType_Hidden), WithInputValue(value))
}

var _ = templruntime.GeneratedTemplate
