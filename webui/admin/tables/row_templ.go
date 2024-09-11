// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package tables

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "context"

type row struct {
	bodyCells []templ.Component
	classes   templ.CSSClasses
	entityId  string
}

type TableRowOption func(*row)

func (l *row) render() templ.Component {
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
		if len(l.entityId) > 0 {
			ctx = context.WithValue(ctx, "itemId", l.entityId)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, row := range l.bodyCells {
			templ_7745c5c3_Err = row.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func WithEntityId(entityId string) TableRowOption {
	return func(c *row) {
		c.entityId = entityId
	}
}

func WithCells(bodyCells ...templ.Component) TableRowOption {
	return func(c *row) {
		c.bodyCells = bodyCells
	}
}

/*
NewTable creates table for listing entities with many TableOptions
*/
func NewRow(options ...TableRowOption) templ.Component {
	row := &row{}
	for _, opt := range options {
		opt(row)
	}

	return row.render()
}
