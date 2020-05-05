package main

import (
	"fmt"

	"github.com/cuckooq/dictionary/dict"
)

func main() {
	dictionary := dict.Dictionary{
		"first":  "first",
		"second": "second",
	}

	// search
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	definition, err = dictionary.Search("sec")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// add
	err = dictionary.Add("third", "third")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	err = dictionary.Add("third", "third")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	// update
	err = dictionary.Update("third", "third2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	err = dictionary.Update("fourth", "fourth")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	// delete
	var word string = "third"
	_, err = dictionary.Search(word)
	if err == nil {
		dictionary.Delete(word)
		fmt.Println(dictionary)
	} else {
		fmt.Println(err)
	}

	word = "third2"
	_, err = dictionary.Search(word)
	if err == nil {
		dictionary.Delete(word)
		fmt.Println(dictionary)
	} else {
		fmt.Println(err)
	}
}
