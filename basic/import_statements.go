// import_statements.go
// Package basic demonstrates various Go import patterns, aliasing, blank imports,
package basic

import (
	// Standard library imports (grouped)
	"fmt"
	"net/http"

	// Aliased imports for clarity or to avoid name conflicts
	console "fmt"
	httpclient "net/http"

	// Blank import: execute init(), but don’t reference the package directly
	_ "net/http/pprof"
)

// ImportExamples shows different import usages
func ImportExamples() {
	// Using unaliased fmt
	fmt.Println("Standard fmt package")

	// Using aliased console
	console.Println("\nAliased fmt as console")

	// Using http package directly
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("\nError via http:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("\nHTTP Status via http:", resp.Status)

	// Using aliased httpclient
	resp2, err2 := httpclient.Get("https://jsonplaceholder.typicode.com/users/1")
	if err2 != nil {
		console.Println("\nError via httpclient:", err2)
		return
	}
	defer resp2.Body.Close()
	console.Println("\nHTTP Status via httpclient:", resp2.Status)

	// Blank import example: pprof init() registered handlers for profiling,
	// accessible via http.DefaultServeMux
	http.DefaultServeMux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)
}
