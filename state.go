package state

import (
	"log"
	"reflect"
	"sync"
)

// State represents state structure
type State struct {
	data reflect.Value

	mapType reflect.Type
	mu      sync.Mutex
}

// New constructs state
func New(keyType, valueType reflect.Type) *State {
	mapType := reflect.MapOf(keyType, valueType)

	return &State{
		data:    reflect.MakeMap(mapType),
		mapType: mapType,
	}
}

// Set adds new key-value pair to state
func (s *State) Set(key, value reflect.Value) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data.SetMapIndex(key, value)

	log.Printf("Data is: %+v", s.data)
}
