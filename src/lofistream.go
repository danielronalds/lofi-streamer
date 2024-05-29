package main

import (
	"encoding/json"
	"os"
)

type LofiStream struct {
	Name    string
	VideoId string
}

func newLofiStream(name string, videoid string) LofiStream {
	return LofiStream{
		Name:    name,
		VideoId: videoid,
	}
}

// Reads JSON from the given file for an array of lofistreams
func readStreamsFromJSON(filename string) ([]LofiStream, error) {
	var streams []LofiStream

	file, err := os.ReadFile(filename)

	err = json.Unmarshal(file, &streams)

	if err != nil {
		return nil, err
	}

	return streams, nil
}
