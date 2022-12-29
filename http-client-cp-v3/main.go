package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ClientGet() ([]Animechan, error) {
	client := http.Client{} //dial

	r, err := client.Get("https://animechan.vercel.app/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var anime []Animechan
	//decode
	err = json.Unmarshal(body, &anime)
	if err != nil {
		panic(err)
	}

	return anime, nil // TODO: replace this
}

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
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
	//encode
	rBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(rBody)

	r, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	//Handle Error
	if err != nil {
		return Postman{}, err
	}
	//baca
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Postman{}, err
	}

	var post Postman
	//decode
	err = json.Unmarshal(body, &post)
	if err != nil {
		return Postman{}, err
	}

	return post, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
