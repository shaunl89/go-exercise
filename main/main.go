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

func main() {

	// terminal prompt for user to enter text and get input text
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your text: ")
	str, _ := reader.ReadString('\n')

	// process the input to remove . , 's elements in text
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
	n := map[int][]string{} // n is type map
	var a []int // type slice
	for k, v := range freq {
		n[v] = append(n[v], k)
	}

	for k := range n {
		a = append(a, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	// if top 3 words exist, print them out if not print everything
	if len(a) > 3 {
		for _, k := range a[0:2] {
			for _, s := range n[k] {
				fmt.Printf("%s, %d times\n", s, k)
			}
		}
	} else {
		for _, k := range a {
			for _, s := range n[k] {
				fmt.Printf("%s, %d times\n", s, k)
			}
		}
	}
}
