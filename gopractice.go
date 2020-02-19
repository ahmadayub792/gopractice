package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=peshawar&appid=83a404e36a0ffc57a902ebfd8f50480b")
	if err != nil {
		log.Fatalln(err)
	}

	result := json:

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}