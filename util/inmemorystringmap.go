package util

import "errors"

// InMemoryStringMap is a test implementation of StringMap.
type InMemoryStringMap struct {
	memory map[string]string
}

// NewInMemoryStringMap works as advertised.
func NewInMemoryStringMap() *InMemoryStringMap {
	return &InMemoryStringMap{
		memory: map[string]string{},
	}
}

// Has tests for the presence of key.
func (m *InMemoryStringMap) Has(key string) (bool, error) {
	_, ok := m.memory[key]
	return ok, nil
}

// Get gets the value of the key.
func (m *InMemoryStringMap) Get(key string) (string, error) {
	val, ok := m.memory[key]
	if !ok {
		return "", errors.New(key + " not in map")
	}
	return val, nil
}

// Set sets the key.
func (m *InMemoryStringMap) Set(key, value string) error {
	m.memory[key] = value
	return nil
}

// Delete removes the key.
func (m *InMemoryStringMap) Delete(key string) error {
	_, ok := m.memory[key]
	if !ok {
		return errors.New("Tried to delete missing " + key)
	}
	delete(m.memory, key)
	return nil
}

// GetAll returns the map
func (m *InMemoryStringMap) GetAll() (map[string]string, error) {
	return m.memory, nil
}
