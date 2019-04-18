package observer

import (
	"reflect"
	"sync"
)

type Observer interface {
	Update(o IObservable, arg interface{})
}

////////////////////////////////////////////////

type IObservable interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	DeleteObservers()
	NotifyObservers(arg interface{})
	SetChanged()
	ClearChanged()
	HasChanged() bool
	CountObservers() int
}

type Observable struct {
	changed bool
	obs map[Observer]bool
	lock *sync.RWMutex
	spec IObservable
}

func NewObservable(spec IObservable) *Observable {
	return &Observable{
		changed: false,
		obs: make(map[Observer]bool),
		lock: new(sync.RWMutex),
		spec: spec,
	}
}

func (o *Observable) AddObserver(observer Observer) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if observer == nil || reflect.ValueOf(observer).IsNil() {
		panic("observer can't be nil")
	}
	if _, ok := o.obs[observer]; !ok {
		o.obs[observer] = false
	}
}

func (o *Observable) DeleteObserver(observer Observer) {
	o.lock.Lock()
	defer o.lock.Unlock()
	delete(o.obs, observer)
}

func (o *Observable) DeleteObservers() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.obs = make(map[Observer]bool)
}

func (o *Observable) NotifyObservers(arg interface{}) {
	o.lock.RLock()
	l := o.CountObservers()
	arrLocal := make([]Observer, l)
	if !o.changed {
		return
	}
	index := 0
	for k := range o.obs {
		arrLocal[index] = k
		index++
	}
	o.changed = false
	o.lock.RUnlock()
	for i:=0; i<l; i++ {
		arrLocal[i].Update(o.spec, arg)
	}
}

func (o *Observable)SetChanged() {
	o.lock.RLock()
	defer o.lock.RUnlock()
	o.changed = true
}

func (o *Observable)ClearChanged(){
	o.lock.RLock()
	defer o.lock.RUnlock()
	o.changed = false
}

func (o *Observable)HasChanged() bool {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.changed
}

func (o *Observable)CountObservers() int {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return len(o.obs)
}