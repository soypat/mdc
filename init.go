package mdc

import (
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
)

/* MDC boot */

const (
	_MDC_VERSION                    = "4.0.0"
	baseStylesheetURL               = "https://unpkg.com/material-components-web@" + _MDC_VERSION + "/dist/material-components-web.min.css"
	robotoMonoStylesheetURL         = "https://fonts.googleapis.com/css?family=Roboto+Mono"
	robotoMonoWeightedStylesheetURL = "https://fonts.googleapis.com/css?family=Roboto+Mono"
	iconsURL                        = "https://fonts.googleapis.com/icon?family=Material+Icons"
)

func AddDefaultStyles() {
	vecty.AddStylesheet(baseStylesheetURL)
	vecty.AddStylesheet(robotoMonoStylesheetURL)
	vecty.AddStylesheet(robotoMonoWeightedStylesheetURL)
	vecty.AddStylesheet(iconsURL)
}

func SetDefaultViewport() {
	meta := js.Global().Get("document").Call("createElement", "meta")
	meta.Set("name", "viewport")
	meta.Set("content", "width=device-width, initial-scale=1")
	js.Global().Get("document").Get("head").Call("appendChild", meta)
}

func AddDefaultScripts() {
	addScript("https://unpkg.com/material-components-web@"+_MDC_VERSION+
		"/dist/material-components-web.min.js", "mdc")
}

func addScript(url string, objName string) {
	script := js.Global().Get("document").Call("createElement", "script")
	script.Set("src", url)
	js.Global().Get("document").Get("head").Call("appendChild", script)
	count := 0
	for {
		count++
		time.Sleep(25 * time.Millisecond)
		if jsObject := js.Global().Get(objName); !jsObject.IsUndefined() {
			break
		} else if count > 100 {
			panic("could not obtain " + objName + " from URL: " + url)
		}
	}
}
