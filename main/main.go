package main

import (
	"fmt"
	"sort"
	"strings"
	"bufio"
	"os"
)

// prompt to input string
// get string from prompt
// remove . , 's and convert to lower case
// split string into array
// store each unique word into object(map?) as key and occurences as value
// sort and display top 10 words

// sample text: "hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time. hello WORLD where hello world World today's art   thou today in this place and time."

func main() {

	// terminal prompt for user to enter text and get input text
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your text: ")
	str, _ := reader.ReadString('\n')
	fmt.Print(str)

	// process input to remove unnecessary elements
	var strLC = strings.ToLower(str)
	var strNoFS = strings.Replace(strLC, ".", "", -1)
	var strNoComma = strings.Replace(strNoFS, ",", "", -1)
	var strNoPlural = strings.Replace(strNoComma, "'s", "", -1)
	var strNoSpaces = strings.Join(strings.Fields(strNoPlural), " ")

	var strArr = strings.Split(strNoSpaces, " ")
	freq := make(map[string]int)

	// storing each element into a map
	for i := range strArr {
		var word = strArr[i]
		if freq[word] >= 1 {
			freq[word]++
		} else {
			freq[word] = 1
		}
	}

	// sorting
	n := map[int][]string{}
	var a []int
	for k, v := range freq {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
			// TODO: work with returning top 3 words
		}
	}
	// TODO: Readme
	// TODO: Github
}
