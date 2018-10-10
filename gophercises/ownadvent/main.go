package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// const storyFile = "gopher.json"

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Gopher Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Story}}
        <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`

type Story map[string]*StoryArc

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func promptForChoice(o []Option) (int, error) {
	var i int
	fmt.Scan(&i)

	if !validateChoice(i, o) {
		fmt.Println("false")
		return 0, fmt.Errorf("not a validate input")
	}
	return i - 1, nil
}

func validateChoice(i int, options []Option) bool {
	if i > 0 && i <= len(options) {
		return true
	}
	return false
}

func getStoryArc(sa *StoryArc) (string, error) {
	fmt.Printf("Title: %s\n", sa.Title)
	for _, story := range sa.Story {
		fmt.Println(story)
		fmt.Println()
	}
	if len(sa.Options) == 0 {
		return "", nil
	}
	fmt.Println("What do you choose?")
	for e, option := range sa.Options {
		fmt.Printf("%d %s:\n", e+1, option.Arc)
		fmt.Printf("%s\n", option.Text)
	}

	nc, err := promptForChoice(sa.Options)
	if err != nil {
		return "", err
	}
	nextStoryArc := sa.Options[nc].Arc
	return nextStoryArc, nil
}

// CLIStory runs a choose your own adventure story via the command line.
func CLIStory(story Story) {
	var err error
	nextStoryArc := "intro"
	fmt.Printf("Choose your own adventure with %s\n\n", story[nextStoryArc].Title)

	for {
		nextStoryArc, err = getStoryArc(story[nextStoryArc])

		if err != nil {
			log.Fatal(err)
		}
		if nextStoryArc == "" {
			break
		}
	}

	fmt.Println("The end")
}

func NewHandler(s Story) http.Handler {

	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func main() {
	var story Story

	advType := flag.String("type", "web", "web or cli")
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	if !containsString([]string{"web", "cli"}, *advType) {
		msg := fmt.Sprintf("do not recognize adventure type %s.  expecting web or cli", *advType)
		log.Fatal(msg)

	}

	fmt.Printf("Using the story %s.\n", *filename)

	sf, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	defer sf.Close()

	sb, err := ioutil.ReadAll(sf)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(sb, &story)
	if err != nil {
		log.Fatal(err)
	}
	if *advType == "cli" {
		CLIStory(story)
	} else {

		h := NewHandler(story)
		fmt.Printf("Starting the server on port %d\n", *port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	}
}

func containsString(strSlice []string, searchStr string) bool {
	for _, value := range strSlice {
		if value == searchStr {
			return true
		}
	}
	return false
}
