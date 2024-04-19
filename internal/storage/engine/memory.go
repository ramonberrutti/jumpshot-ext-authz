package engine

import (
	"context"
	"errors"
	"sync"
)

func NewMemory() Engine {
	return &memory{
		keys: make(map[string]*memoryData),
	}
}

type memory struct {
	keys map[string]*memoryData
	m    sync.RWMutex
}

type memoryData struct {
	data    []byte
	version uint64
}

func (m *memory) Set(ctx context.Context, key string, data []byte) error {
	m.m.Lock()
	defer m.m.Unlock()

	_, ok := m.keys[key]
	if ok {
		return ErrNoAction
	}

	m.keys[key] = &memoryData{
		data:    data,
		version: 1,
	}
	return nil
}

func (m *memory) Get(ctx context.Context, key string) ([]byte, error) {
	m.m.RLock()
	defer m.m.RUnlock()

	value, ok := m.keys[key]
	if !ok {
		return nil, ErrNotFound
	}

	return value.data, nil
}

func (m *memory) Update(
	ctx context.Context,
	key string,
	f UpdateFunc,
) error {
	// Simulate a real engine with this two locks.
	m.m.RLock()
	value, ok := m.keys[key]
	if !ok {
		m.m.RUnlock()
		return ErrNotFound
	}
	d := value.data
	v := value.version
	m.m.RUnlock()

	newData, err := f(ctx, d)
	if err != nil {
		return err
	}

	// Simulate a real engine with this two locks.
	m.m.Lock()
	defer m.m.Unlock()
	oldValue, ok := m.keys[key]
	if !ok {
		return ErrNotFound
	}

	if v != oldValue.version {
		return errors.New("unable to update")
	}

	m.keys[key] = &memoryData{
		data:    newData,
		version: v + 1,
	}
	return nil
}

func (m *memory) Delete(ctx context.Context, key string) (bool, error) {
	m.m.Lock()
	defer m.m.Unlock()
	_, ok := m.keys[key]
	if !ok {
		return false, ErrNotFound
	}

	delete(m.keys, key)
	return true, nil
}
