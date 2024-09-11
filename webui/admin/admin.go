// This package contains mandatory structures and interfaces for admin ui kit
package admin

import (
	"github.com/a-h/templ"
)

// AdminComponent it`s base component interface
type AdminComponent interface {
	Build() templ.Component
}

// ListResponseMeta defaults response meta for
type ListResponseMeta struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
