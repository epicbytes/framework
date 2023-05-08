package framework

import (
	"context"
	"github.com/epicbytes/framework/storage/mongodb"
	"net/http"
	"sync"
)

// BaseTwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type BaseTwirpServer interface {
	http.Handler

	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// google.golang.org/protobuf/types/descriptorpb.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)

	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string

	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route Twirp requests.
	// The path prefix is in the form: "/<prefix>/<package>.<Service>/"
	// that is, everything in a Twirp route except for the <Method> at the end.
	PathPrefix() string
}

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
