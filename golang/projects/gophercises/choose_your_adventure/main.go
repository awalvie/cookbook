package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (arc Arc) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println("Error while decoding", err)
		return err
	}

	fmt.Println(v)
	return nil
}

func main() {
	// read file
	data, err := ioutil.ReadFile("./story.json")
	if err != nil {
		fmt.Println(err)
	}

	var arcs []Arc

}
