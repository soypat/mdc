package mdc

import (
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
)

/* MDC boot */

// See https://material.io/develop/web/getting-started
const (
	_MDC_VERSION                    = "13.0.0" // use `latest` to always acquire latest css.
	baseStylesheetURL               = "https://unpkg.com/material-components-web@" + _MDC_VERSION + "/dist/material-components-web.min.css"
	baseScriptURL                   = "https://unpkg.com/material-components-web@" + _MDC_VERSION + "/dist/material-components-web.min.js"
	robotoMonoStylesheetURL         = "https://fonts.googleapis.com/css?family=Roboto+Mono"
	robotoMonoWeightedStylesheetURL = "https://fonts.googleapis.com/css?family=Roboto:300,400,500,700"
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
	addScript(baseScriptURL, "mdc")
}

func mdcOK() bool {
	mdc := js.Global().Get("mdc")
	return !mdc.IsNull() && !mdc.IsUndefined()
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

const svgCheckbox = `<svg class="mdc-checkbox__checkmark"
viewBox="0 0 24 24">
<path class="mdc-checkbox__checkmark-path"
   fill="none"
   d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
</svg>
<div class="mdc-checkbox__mixedmark"></div>`

const spaCSS = `// Note: These styles do not account for any paddings/margins that you may need.

body {
  display: flex;
  height: 100vh;
}

.mdc-drawer-app-content {
  flex: auto;
  overflow: auto;
  position: relative;
}

.main-content {
  overflow: auto;
  height: 100%;
}

.app-bar {
  position: absolute;
}

// Only apply this style if below top app bar.
.mdc-top-app-bar {
  z-index: 7;
}`
