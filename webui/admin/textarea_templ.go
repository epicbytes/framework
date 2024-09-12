// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type TextareaModifier string
type TextareaResponsive string

const (
	// TextareaModifier_Bordered adds border to textarea
	TextareaModifier_Bordered TextareaModifier = "textarea-bordered"
	// TextareaModifier_Ghost adds ghost style to textarea
	TextareaModifier_Ghost TextareaModifier = "textarea-ghost"
	// TextareaModifier_Primary adds `primary` color to textarea
	TextareaModifier_Primary TextareaModifier = "textarea-primary"
	// TextareaModifier_Secondary adds `secondary` color to textarea
	TextareaModifier_Secondary TextareaModifier = "textarea-secondary"
	// TextareaModifier_Accent adds `accent` color to textarea
	TextareaModifier_Accent TextareaModifier = "textarea-accent"
	// TextareaModifier_Info adds `info` color to textarea
	TextareaModifier_Info TextareaModifier = "textarea-info"
	// TextareaModifier_Success adds `success` color to textarea
	TextareaModifier_Success TextareaModifier = "textarea-success"
	// TextareaModifier_Warning adds `warning` color to textarea
	TextareaModifier_Warning TextareaModifier = "textarea-warning"
	// TextareaModifier_Error adds `error` color to textarea
	TextareaModifier_Error TextareaModifier = "textarea-error"
)

const (
	// TextareaResponsive_Lg for Large size for textarea
	TextareaResponsive_Lg TextareaResponsive = "textarea-lg"
	// TextareaResponsive_Md for Medium (default) size for textarea
	TextareaResponsive_Md TextareaResponsive = "textarea-md"
	// TextareaResponsive_Sm for Small size for textarea
	TextareaResponsive_Sm TextareaResponsive = "textarea-sm"
	// TextareaResponsive_Xs for Extra small size for textarea
	TextareaResponsive_Xs TextareaResponsive = "textarea-xs"
)

type textarea struct {
	label              string
	name               string
	placeholder        string
	value              string
	classes            templ.CSSClasses
	textareaModifier   []TextareaModifier
	textareaResponsive []TextareaResponsive
	attrs              templ.Attributes
}

type TextareaOption func(*textarea)

func (b *textarea) buildClasses() templ.CSSClasses {
	var classes string

	classes = b.classes.String()

	classes += " " + strings.Join(lo.Map(b.textareaModifier, func(item TextareaModifier, index int) string {
		return string(item)
	}), " ")

	classes += " " + strings.Join(lo.Map(b.textareaResponsive, func(item TextareaResponsive, index int) string {
		return string(item)
	}), " ")

	unifiedClasses := lo.Uniq(strings.Split(classes, " "))
	return templ.Classes(unifiedClasses)
}

func (b *textarea) getError(ctx context.Context) string {
	if values, ok := ctx.Value("errors").(map[string]string); ok {
		return lo.ValueOr(values, b.name, "")
	} else {
		return ""
	}
}

func (b *textarea) render() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"form-control w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(b.label) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"label\"><span class=\"label-text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(b.label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 87, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grow-wrap\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 = []any{"textarea", b.buildClasses(), templ.KV("textarea-error", len(b.getError(ctx)) > 0)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<textarea name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(b.name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 91, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(b.placeholder)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 91, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(b.value)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 91, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" oninput=\"this.parentNode.dataset.value = this.value\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var3).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
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
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(b.value)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 91, Col: 252}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(b.getError(ctx)) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"label\"><span class=\"label-text-alt text-error\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(b.getError(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/textarea.templ`, Line: 95, Col: 61}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func WithTextareaValue(value any) TextareaOption {
	return func(b *textarea) {
		b.value = fmt.Sprintf("%v", value)
	}
}

func WithTextareaLabel(label string) TextareaOption {
	return func(b *textarea) {
		b.label = label
	}
}

func WithTextareaPlaceholder(placeholder string) TextareaOption {
	return func(b *textarea) {
		b.placeholder = placeholder
	}
}

func WithTextareaModifiers(textareaModifier ...TextareaModifier) TextareaOption {
	return func(b *textarea) {
		b.textareaModifier = lo.Uniq(append(b.textareaModifier, textareaModifier...))
	}
}

func WithTextareaResponsive(textareaResponsive ...TextareaResponsive) TextareaOption {
	return func(b *textarea) {
		b.textareaResponsive = lo.Uniq(append(b.textareaResponsive, textareaResponsive...))
	}
}

func WithTextareaAttrs(attrs templ.Attributes) TextareaOption {
	return func(b *textarea) {
		b.attrs = attrs
	}
}

/*
NewTextarea creates input with some props
*/
func NewTextarea(name string, options ...TextareaOption) templ.Component {
	tar := &textarea{
		name:             name,
		attrs:            templ.Attributes{"rows": 1},
		textareaModifier: []TextareaModifier{TextareaModifier_Bordered},
	}

	for _, opt := range options {
		opt(tar)
	}

	return tar.render()
}

/*
NewTextareaSm creates input smaller with some props
Textareas allow the user to input values.
*/
func NewTextareaSm(name string, options ...TextareaOption) templ.Component {
	return NewTextarea(name, append(options, WithTextareaResponsive(TextareaResponsive_Sm))...)
}