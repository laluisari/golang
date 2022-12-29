package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}

	resp, err := client.Get("https://animechan.vercel.app/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var anime []Animechan
	err = json.Unmarshal(body, &anime)
	if err != nil {
		panic(err)
	}

	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	return anime, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	//Handle Error
	if err != nil {
		return Postman{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	var post Postman
	err = json.Unmarshal(body, &post)
	if err != nil {
		return Postman{}, err
	}

	// Hit API https://postman-echo.com/post with method POST:
	return post, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
