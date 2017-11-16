package main

import (
	"fmt"
)

type element struct {
	x int
	y int
	User string
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string `json:"user"`
	Message  string  `json:"message"`
	Location Location `json:"location"`
}

func main() {
	//Homework 1, given a map from string to bool, how to iterate over it and print out all the key-values?
	m := make(map[string]bool)
	m["andy"] = true
	m["monica"] = false
	for k, v := range m {
		fmt.Println(k, v)
	}

}
