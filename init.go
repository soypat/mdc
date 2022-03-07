package mdc

import (
	"io"
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
	"github.com/soypat/gwasm"
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

func AddDefaultScripts(timeout time.Duration) {
	gwasm.AddScript(baseScriptURL, "mdc", timeout)
}

func mdcOK() bool {
	mdc := js.Global().Get("mdc")
	return !mdc.IsNull() && !mdc.IsUndefined()
}

// Console returns the browser console as an io.Writer which one
// can use to set log output.
func Console() io.Writer {
	return jsWriter{
		Value: js.Global().Get("console"),
		fname: "log",
	}
}

type jsWriter struct {
	js.Value
	fname string
}

func (j jsWriter) Write(b []byte) (int, error) {
	j.Call(j.fname, string(b))
	return len(b), nil
}

const svgCheckbox = `<svg class="mdc-checkbox__checkmark"
viewBox="0 0 24 24">
<path class="mdc-checkbox__checkmark-path"
   fill="none"
   d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
</svg>
<div class="mdc-checkbox__mixedmark"></div>`

const spaCSS = `/* Note: These styles do not account for any paddings/margins that you may need. */

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

/* Only apply this style if below top app bar. */
.mdc-top-app-bar {
  z-index: 7;
}`
