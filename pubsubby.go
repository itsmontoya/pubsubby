package pubsubby

import (
	"sync"

	"github.com/cheekybits/genny/generic"
	"github.com/itsmontoya/pubsubby/utilities"
)

// newPubsubby will return a new instance of pubsubby
// Note: This is private because this library is intended to be generated into other libraries
func newPubsubby() *pubsubby {
	var p pubsubby
	p.psm = make(map[Key]*pubsub)
	return &p
}

// pubsubby manages a set of pubsub's
type pubsubby struct {
	mux sync.RWMutex
	psm map[Key]*pubsub
}

// get will attempt to get a pubsub for a given key
// Note: This function is thread-safe, locking does not need to be handled elsewhere
func (p *pubsubby) get(key Key) (ps *pubsub, ok bool) {
	p.mux.RLock()
	ps, ok = p.psm[key]
	p.mux.RUnlock()
	return
}

// create will create a pubsub for a given key if the pubsub does not yet exist
// Note: This function is thread-safe, locking does not need to be handled elsewhere
func (p *pubsubby) create(key Key) (ps *pubsub) {
	var ok bool
	// Attempt to get the value first, this will allow us to avoid a write-lock if the value exists
	if ps, ok = p.get(key); ok {
		// Fast-track successful, return early
		return
	}

	p.mux.Lock()
	// Check if the value still does not exist (in case it was created before new lock was acquired)
	if ps, ok = p.psm[key]; !ok {
		ps = &pubsub{}
		p.psm[key] = ps
	}
	p.mux.Unlock()
	return
}

// Subscribe will add a subscriber to the functions list for a matching pubsub key
func (p *pubsubby) Subscribe(key Key, fn SubFn) {
	ps := p.create(key)
	ps.Subscribe(fn)
}

// Publish will publish a value to the subscribers for a matching pubsub key
func (p *pubsubby) Publish(key Key, val Value) {
	ps := p.create(key)
	ps.Publish(key, val)
}

// pubsub is a pubsub item
type pubsub struct {
	mux sync.RWMutex
	fns []SubFn
}

func (p *pubsub) pop(i int) {
	p.fns = append(p.fns[:i], p.fns[i+1:]...)
}

// Subscribe will add a subscriber to the functions list
func (p *pubsub) Subscribe(fn SubFn) {
	p.mux.Lock()
	p.fns = append(p.fns, fn)
	p.mux.Unlock()
}

// Publish will publish a value to the subscribers
func (p *pubsub) Publish(key Key, val Value) {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Iterate through all the subscribers
	for i, fn := range p.fns {
		if fn(key, val) {
			// Function's end variable returned as true, pop the function from the subscribers list
			p.pop(i)
		}
	}
}

// Len is used to determine the length of the subscribers
func (p *pubsub) Len() (n int) {
	p.mux.RLock()
	n = len(p.fns)
	p.mux.RUnlock()
	return
}

// List is for debugging purposes, will allow to peek at the current subscibers
func (p *pubsub) List() (fis []utilities.FuncInfo) {
	p.mux.RLock()
	defer p.mux.RUnlock()

	for _, fn := range p.fns {
		fis = append(fis, utilities.NewFuncInfo(fn))
	}

	return
}

// SubFn will take a value and return an "end" boolean
type SubFn func(key Key, val Value) (end bool)

// Key is the key type
type Key generic.Type

// Value is the value type
type Value generic.Type
