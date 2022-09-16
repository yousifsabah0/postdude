package postdude

import (
	"fmt"
	"net/http"
	"strings"
)

// Builder acts like helper to setup the request.
type Builder struct {
	prefix string
	strings.Builder
}

// Println prints a new line to the console.
func (b *Builder) Println() {
	b.WriteString("\n")
}

// Printf prints a formatted given string to the console.
func (b *Builder) Printf(s string, a ...any) {
	b.WriteString(fmt.Sprintf("%v %v\n", b.prefix, fmt.Sprintf(s, a...)))
}

// WriteHeader writes the current request headers to the console.
func (b *Builder) WriterHeader(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			b.Printf("%v : %v", key, value)
		}
	}
}
