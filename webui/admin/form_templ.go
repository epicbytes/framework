// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"fmt"
	"github.com/epicbytes/framework/webui/shared"
	"github.com/samber/lo"
	"net/url"
)

type FormType string
type FormMethod string

const (
	FormType_List   FormType = ""
	FormType_Create FormType = "create"
	FormType_Update FormType = "update"
)

const (
	FormMethod_Get    FormMethod = "GET"
	FormMethod_Post   FormMethod = "POST"
	FormMethod_Put    FormMethod = "PUT"
	FormMethod_Delete FormMethod = "DELETE"
)

type form struct {
	urlPrefix  string
	entityId   string
	entityName string
	formId     string
	formType   FormType
	formMethod FormMethod
	attrs      templ.Attributes
	classes    templ.CSSClasses
}

func (f *form) makeEntityUrl(ctx context.Context) string {
	url, _ := url.JoinPath("/", shared.GetEntityUrlPrefix(ctx), shared.GetEntityName(ctx), string(f.formType), fmt.Sprintf("%v", f.entityId))
	return url
}

type FormOption func(*form)

func WithFormType(formType FormType) FormOption {
	return func(f *form) {
		f.formType = formType
	}
}

func WithFormMethod(formMethod FormMethod) FormOption {
	return func(f *form) {
		f.formMethod = formMethod
	}
}

func WithEntityId(entityId string) FormOption {
	return func(f *form) {
		f.entityId = entityId
	}
}

func WithEntityName(entityName string) FormOption {
	return func(f *form) {
		f.entityName = entityName
	}
}

func WithFormAttrs(attrs templ.Attributes) FormOption {
	return func(f *form) {
		f.attrs = attrs
	}
}

func WithFormUrlPrefix(urlPrefix string) FormOption {
	return func(f *form) {
		f.urlPrefix = urlPrefix
	}
}

func WithFormId(formId string) FormOption {
	return func(f *form) {
		f.formId = formId
	}
}

func (f *form) render() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL(f.makeEntityUrl(ctx))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(f.formId) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(f.formId)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 91, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" method=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(lo.Ternary(f.formMethod == FormMethod_Get, "GET", "POST"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 93, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" enctype=\"multipart/form-data\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if f.formMethod == FormMethod_Post {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(f.makeEntityUrl(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 96, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if f.formMethod == FormMethod_Put {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-put=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(f.makeEntityUrl(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 99, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if f.formMethod == FormMethod_Delete {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-delete=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(f.makeEntityUrl(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 102, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if f.formMethod == FormMethod_Get {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(f.makeEntityUrl(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `admin/form.templ`, Line: 105, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func NewForm(options ...FormOption) templ.Component {
	frm := &form{}

	for _, opt := range options {
		opt(frm)
	}

	return frm.render()
}

func NewCreateForm(options ...FormOption) templ.Component {
	return NewForm(append(options, WithFormType(FormType_Create), WithFormMethod(FormMethod_Post), WithFormAttrs(templ.Attributes{"hx-target": "this", "hx-swap": "outerHTML"}))...)
}

func NewUpdateForm(entity string, entityId string, urlPrefix string, options ...FormOption) templ.Component {
	return NewForm(append(options, WithFormType(FormType_Update), WithFormMethod(FormMethod_Post), WithEntityId(entityId), WithFormUrlPrefix(urlPrefix))...)
}
