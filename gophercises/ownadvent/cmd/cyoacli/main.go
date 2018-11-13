package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mudzi42/playground/gophercises/ownadvent"
)

func main() {
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

	CLIStory(story)
}

func promptForChoice(o []cyoa.Option) (int, error) {
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatal(err)
	}

	if !validateChoice(i, o) {
		fmt.Println("false")
		return 0, fmt.Errorf("not a validate input")
	}
	return i - 1, nil
}

func validateChoice(i int, options []cyoa.Option) bool {
	if i > 0 && i <= len(options) {
		return true
	}
	return false
}

func getStoryArc(sa *cyoa.StoryArc) (string, error) {
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
func CLIStory(story cyoa.Story) {
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
