package main

import (
	"fmt"

	"github.com/hwalim/go-project/dict/mydict"
)

// dictionary search
// func main() {
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	definition, err := dictionary.Search("first")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(definition)
// 	}
// }

// dictionary add
// func main() {
// 	dictionary := mydict.Dictionary{"first": " First word"}
// 	word := "second"
// 	def := "second word"
// 	err := dictionary.Add(word, def)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defi, _ := dictionary.Search(word)
// 	fmt.Println(defi)

// }

// update word
// func main() {
// 	dictionary := mydict.Dictionary{}
// 	word := "first"
// 	def := "first word"
// 	dictionary.Add(word, def)
// 	err := dictionary.Update(word, "second word")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	value, _ := dictionary.Search(word)
// 	fmt.Println(value)
// }

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	word := "second"
	err := dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)

}
