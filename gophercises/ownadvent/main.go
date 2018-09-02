package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const storyFile = "gopher.json"

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

func main() {
	fmt.Printf("Choose your own adventure with %s\n\n", storyFile)

	sf, err := os.Open(storyFile)
	if err != nil {
		fmt.Println(err)
	}
	defer sf.Close()

	sb, err := ioutil.ReadAll(sf)
	if err != nil {
		fmt.Println(err)
	}

	storyArcs := make(map[string]*StoryArc)

	err = json.Unmarshal(sb, &storyArcs)
	if err != nil {
		log.Fatal(err)
	}

	nextStoryArc := "intro"
	for {
		nextStoryArc, err = getStoryArc(storyArcs[nextStoryArc])
		if err != nil {
			log.Fatal(err)
		}
		if nextStoryArc == "" {
			break
		}
	}

	fmt.Println("The end")
}
