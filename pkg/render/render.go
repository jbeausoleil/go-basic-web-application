package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config

	tc, err := CreateTemplateCache() // builds template cache during each page visit
	if err != nil {
		log.Fatal(err)
	}

	// tmpl is passed from the handler functions which will match myCache[name]
	t, ok := tc[tmpl] // ok set to true if tmpl exists
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil) // appends the contents of t to the buffer buf, but do not pass it data
	_, err = buf.WriteTo(w) // write data to the ResponseWriter until the buffer is drained
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// setup store for template name and address of template in memory
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl") // create a slice of files that are found at the path specified
	if err != nil {
		return myCache, err
	}

	// loop through slice of pages
	for _, page := range pages {
		name := filepath.Base(page)                                     // extract template name from path
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // create new template set from parsing path to tmpl
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl") // create a slice of files that are found at the path specified
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // parse template definitions for all files found
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
