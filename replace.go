//replace non-alphanumerics with dash, spaced or nothing/ remove non alphanumerics 

package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString("#Go#Python$Php&Perl@@", " ")
	//newStr := reg.ReplaceAllString("#Go#Python$Php&Perl@@", "")//replace with nothing; remove
	fmt.Println(newStr)
}