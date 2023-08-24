package framework

import (
	"context"
	"github.com/epicbytes/framework/storage/mongodb"
	"sync"
)

type Service interface {
	GetModel(name mongodb.CollectionName) mongodb.Model
	GetStore(name mongodb.CollectionName) *sync.Map
}

type BaseListener struct {
	Framework Framework
}

func (s *BaseListener) GetModel(name mongodb.CollectionName) mongodb.Model {
	return s.Framework.GetModel(name)
}

func (s *BaseListener) GetStore(name mongodb.CollectionName) *sync.Map {
	return s.Framework.GetStore(name)
}

type BaseGateway interface {
	SetFramework(frm ForGateway)
	Start(ctx context.Context)
}
