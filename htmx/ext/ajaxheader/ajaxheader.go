// package ajaxheader adds the X-Requested-With header to requests with the value “XMLHttpRequest”.
package ajaxheader

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension adds the X-Requested-With header to requests with the value “XMLHttpRequest”.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@2.0.7/dist/ext/ajax-header.js"></script>
//
// # Usage
//
//	<body { hx.Ext(ajaxheader.Extension)... } >
//
// Extension: [ajax-header]
//
// [ajax-header]: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/ajax-header/README.md
const Extension htmx.Extension = "ajax-header"
