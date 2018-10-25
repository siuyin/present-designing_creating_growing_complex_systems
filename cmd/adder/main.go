// 10 OMIT
package main

// 20 OMIT

import (
	"html/template"
	"log"
	"net/http"

	"github.com/siuyin/present-designing_creating_growing_complex_systems/adder"
)

// 30 OMIT
func main() {
	log.Println("adder starting")
	http.HandleFunc("/", rootHandler)
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static"))))

	// Adder area starts
	http.HandleFunc("/adder", adder.AdderHandler)
	// Adder area ends

	http.ListenAndServe(":8080", nil)
}

// 40 OMIT

//50 OMIT
const mainScreen = `
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/static/css/pure-min.css">
  <link rel="stylesheet" href="/static/css/app.css">
  <title>Adder</title>
</head>

<body>
<h1>Adder system</h1>

{{- .AdderForm -}}
<div id="adder-result"></div>

<script src="/static/js/adder.js"></script>
</body>

</html>
`

//60 OMIT

//70 OMIT
var mainTpl *template.Template

func init() {
	mainTpl = template.Must(template.New("main").Parse(mainScreen))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if err := mainTpl.Execute(w, struct {
		AdderForm template.HTML
	}{
		AdderForm: adderForm,
	}); err != nil {
		log.Fatalf("main execute template: %v", err)
	}
}

//80 OMIT
