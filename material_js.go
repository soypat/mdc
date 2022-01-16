package mdc

import "syscall/js"

type namespace string

const (
	nsTooltip namespace = "tooltip"
	nsSlider  namespace = "slider"
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
	el := js.Global().Get("document").Call("getElementById", id)
	if el.IsNull() {
		panic("element of id " + id + " not found")
	}
	return ns.new(obj, el)
}

// Equivalent of
//  value = new mdc.namespace.obj(document.querySelector(selector))
func (ns namespace) newFromQuery(obj, selector string) js.Value {
	el := js.Global().Get("document").Call("querySelector", selector)
	if el.IsNull() {
		panic("no elements found with selector: " + selector)
	}
	return ns.new(obj, el)
}
