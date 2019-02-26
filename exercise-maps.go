package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	sWords := strings.Fields(s)
	sMap := make(map[string]int)
	
	for i:= 0; i < len(sWords); i++ {
	  
    //Increases the count of a word if it's already been stored in the map
    //Otherwise, stores it with a count of '1'
		if sMap[sWords[i]] == 1 {
			sMap[sWords[i]]++
		} else {
			sMap[sWords[i]] = 1
		}
	}
	
	
	return sMap
}

func main() {
	wc.Test(WordCount)
}
