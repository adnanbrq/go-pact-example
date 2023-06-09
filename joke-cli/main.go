package main

import (
	"fmt"
	"io"
	"net/http"
)

type Joke struct {
	joke string `json:"joke"`
}

func getRandomJoke(url string) (Joke, error) {
	res, err := http.Get(url)
	if err != nil {
		return Joke{}, err
	}

	joke, err := io.ReadAll(res.Body)
	if err != nil {
		return Joke{}, err
	}

	return Joke{string(joke)}, nil
}

func main() {
	joke, err := getRandomJoke("http://localhost:3031")
	if err != nil {
		panic(err)
	}

	fmt.Println(joke.joke)
}
