package helpers

import (
	"context"
	"os"
)

var IsDev = os.Getenv("ENVIRONMENT") == "development"

type ContextKey string

const (
	CurrentLanguage  ContextKey = "currentLanguage"
	CurrentPath      ContextKey = "currentPath"
	BuildTimestamp   ContextKey = "buildTimestamp"
	CsrfToken        ContextKey = "csrfToken"
	Token            ContextKey = "token"
	InputName        ContextKey = "inputName"
	InputParentName  ContextKey = "parentName"
	InputParentIndex ContextKey = "parentIndex"
	InputParentIsMap ContextKey = "parentIsMap"
)

func GetStringValueFromCtx(ctx context.Context, key ContextKey) string {
	if len(key) == 0 {
		return ""
	}
	if value, ok := ctx.Value(string(key)).(string); ok {
		return value
	}
	return ""
}

func GetIntValueFromCtx(ctx context.Context, key ContextKey) *int {
	if value, ok := ctx.Value(string(key)).(int); ok {
		return &value
	}
	var value int
	return &value
}

func GetBoolValueFromCtx(ctx context.Context, key ContextKey) bool {
	if len(key) == 0 {
		return false
	}
	if value, ok := ctx.Value(string(key)).(bool); ok {
		return value
	}
	return false
}

func GetHost() string {
	return os.Getenv("MAIN_HOST")
}
