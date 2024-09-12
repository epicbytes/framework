package grpc

import "context"

type Handler = func(ctx context.Context, args interface{})

type ServiceHookIdentifier string

func (s ServiceHookIdentifier) String() string {
	return string(s)
}

type ServiceHooks struct {
	handlers map[ServiceHookIdentifier][]Handler
}

func (s *ServiceHooks) On(eventName ServiceHookIdentifier, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[ServiceHookIdentifier][]Handler)
	}
	if _, ok := s.handlers[eventName]; !ok {
		s.handlers[eventName] = make([]Handler, 0)
	}
	s.handlers[eventName] = append(s.handlers[eventName], handler)
}

func (s *ServiceHooks) Trigger(ctx context.Context, eventName ServiceHookIdentifier, args interface{}) {
	if s.handlers == nil {
		return
	}
	if handlers, ok := s.handlers[eventName]; ok {
		for _, handler := range handlers {
			handler(ctx, args)
		}
	}
}
