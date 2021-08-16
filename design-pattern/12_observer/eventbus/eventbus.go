package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type IBus interface {
	Subscribe(string, interface{}) error
	Publish(string, ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("handler must be a function.")
	}
	handlers, ok := bus.handlers[topic]
	if !ok {
		handlers = []reflect.Value{}
	}
	handlers = append(handlers, v)
	bus.handlers[topic] = handlers
	return nil
}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Printf("No Handlers found in topic %s\n", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	for i := range handlers {
		go handlers[i].Call(params)
	}
}
