package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mudzi42/playground/gophercises/ownadvent"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	fmt.Printf("Using the story %s.\n", *filename)

	sf, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = sf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	story, err := cyoa.JSONStory(sf)
	if err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(
		story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)

	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	mux.Handle("/", cyoa.NewHandler(story))

	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

// Updated chapter parsing function. Technically you don't
// *have* to get the story from the path (it could be a
// header or anything else) but I'm not going to rename this
// since "path" is what we used in the videos.
func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// Slightly altered tempalte to show how this feature works
var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
