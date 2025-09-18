// package removeme allows you to remove an element after a specified interval.
package removeme

import (
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to remove an element after a specified interval.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@2.0.7/dist/ext/remove-me.js"></script>
//
// Extension: [remove-me]
//
// [remove-me]: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/remove-me/README.md
const Extension htmx.Extension = "remove-me"

// RemoveMe removes the element after the specified interval.
//
// Extension: [remove-me]
//
// [remove-me]: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/remove-me/README.md
func RemoveMe[T any](hx htmx.HX[T], after time.Duration) T {
	return hx.Attr("remove-me", after.String())
}
