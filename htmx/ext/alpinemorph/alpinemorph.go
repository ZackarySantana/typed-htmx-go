// package alpinemorph allows you to use the Alpine.js lightweight [morph plugin] as the swapping mechanism in htmx which is necessary to retain Alpine state when you have entire Alpine components swapped by htmx.
//
// [morph plugin]: https://alpinejs.dev/plugins/morph
package alpinemorph

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension allows you to use the Alpine.js lightweight [morph plugin] as the swapping mechanism in htmx which is necessary to retain Alpine state when you have entire Alpine components swapped by htmx.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@2.0.7/dist/ext/alpine-morph.js"></script>
//
// Extension: [alpine-morph]
//
// [alpine-morph]: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/alpine-morph/README.md
// [morph plugin]: https://alpinejs.dev/plugins/morph
const Extension htmx.Extension = "alpine-morph"
