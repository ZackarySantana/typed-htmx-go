// package debug logs all htmx events for the element it is on, either through the console.debug function or through the console.log function with a DEBUG: prefix.
package debug

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension logs all htmx events for the element it is on, either through the console.debug function or through the console.log function with a DEBUG: prefix.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@2.0.7/dist/ext/debug.js"></script>
//
// # Usage
//
//	<button { hx.Ext(debug.Extension)... } >
//
// Extension: [debug]
//
// [debug]: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/debug/README.md
const Extension htmx.Extension = "debug"
