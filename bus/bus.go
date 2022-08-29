package bus

import "sync"

type LocalEventBus struct {
}

type Publisher interface {
}

type Listener interface {
	sub()
}

type EventBus interface {
	AddListener(topic string, listener Listener) (bool, error)
}

type bus struct {
	ch   chan interface{}
	lock sync.Mutex
}

func (b *LocalEventBus) New() error {
	return nil
}
