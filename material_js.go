package mdc

import (
	"errors"
	"strconv"
	"syscall/js"
)

var globalHandlers = newHandlerStore()

var (
	errHandlerAlreadyRegistered = errors.New("tried to register an existing/empty id for a javascript handle to a MDC component. " +
		"Make sure IDs are unique or that you are not rendering single page components (i.e. Leftbar, Navbar) multiple times")
)

type JSComponent interface {
	id() string
}

// Handler returns the underlying javascript object which controls
// the rendering of c.
func Handler(c JSComponent) js.Value {
	return globalHandlers.getID(c.id())
}

// destroyHandler calls the finalizer of the javascript handler
// and unregisters the id.
func destroyHandler(c JSComponent) {
	id := c.id()
	globalHandlers.getID(id).Call("destroy")
	globalHandlers.unregisterID(id)
}

type handlerStore struct {
	// ID registered handlers.
	id map[string]js.Value
	// Automatic ID assignation scheme.
	auto int
}

func newHandlerStore() handlerStore {
	return handlerStore{
		id: make(map[string]js.Value),
	}
}

func (hs *handlerStore) getUnique() (id string) {
	for {
		hs.auto++
		id = "umdc" + strconv.Itoa(hs.auto)
		if hs.isFree(id) {
			return id
		}
	}
}

func (hs handlerStore) isFree(id string) bool {
	if id == "" {
		return false
	}
	_, present := hs.id[id]
	return !present
}

func (hs handlerStore) registerID(id string, handler js.Value) error {
	if !hs.isFree(id) {
		panic(errHandlerAlreadyRegistered.Error() + " concerning id: \"" + id + "\"") // TODO(soypat): Remove this panic when ready for production.
		return errHandlerAlreadyRegistered
	}
	hs.id[id] = handler
	return nil
}

func (hs handlerStore) getID(id string) js.Value {
	return hs.id[id]
}

func (hs handlerStore) unregisterID(id string) {
	delete(hs.id, id)
}

type namespace string

const (
	nsTooltip namespace = "tooltip"
	nsSlider  namespace = "slider"
	nsDrawer  namespace = "drawer"
	nsList    namespace = "list"
	nsMenu    namespace = "menu"
)

func (ns namespace) String() string {
	return string(ns)
}

// Equivalent of
//  value = mdc.namespace
func (ns namespace) get() js.Value {
	if !mdcOK() {
		panic("mdc must be defined to acquire namespace")
	}
	return js.Global().Get("mdc").Get(ns.String())
}

// Equivalent of
//  value = new mdc.namespace.obj(args)
func (ns namespace) new(obj string, args ...interface{}) js.Value {
	v := ns.get().Get(obj)
	if v.IsNull() {
		panic(obj + " is not defined in mdc." + ns.String())
	}
	return v.New(args...)
}

// Equivalent of
//  value = new mdc.namespace.obj(document.getElementById(id))
func (ns namespace) newFromId(obj, id string) js.Value {
	el := docGetByID(id)
	if el.IsNull() {
		panic("element of id " + id + " not found")
	}
	return ns.new(obj, el)
}

// Equivalent of
//  value = new mdc.namespace.obj(document.querySelector(selector))
func (ns namespace) newFromQuery(obj, selector string) js.Value {
	el := docQuery(selector)
	if el.IsNull() {
		panic("no elements found with selector: " + selector)
	}
	return ns.new(obj, el)
}

func docGetByID(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func docQuery(selector string) js.Value {
	return js.Global().Get("document").Call("querySelector", selector)
}
