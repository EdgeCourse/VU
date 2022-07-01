//validate credit card number

/*

((4|5)\d{3}-?\d{4}-?\d{4}-?\d{4}|(4|5)\d{15})|(^(6011)-?\d{4}-?\d{4}-?\d{4}|(6011)-?\d{12})|(^((3\d{3}))-\d{6}-\d{5}|^((3\d{14})))

*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "4111112111311111"
	str2 := "341823205239073"
	str3 := "370751517018351"
	str4 := "4556219835495866"
	str5 := "5019727010303742"
	str6 := "76029244561"
	str7 := "4111-1111-1121-1111"
	str8 := "5610591081019250"
	str9 := "30559309025904"
	str10 := "6011121112111117"

	re := regexp.MustCompile(`((4|5)\d{3}-?\d{4}-?\d{4}-?\d{4}|(4|5)\d{15})|(^(6011)-?\d{4}-?\d{4}-?\d{4}|(6011)-?\d{12})|(^((3\d{3}))-\d{6}-\d{5}|^((3\d{14})))`)

	fmt.Printf("\nCC : %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("CC : %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("CC : %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("CC : %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("CC : %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("CC : %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("CC : %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("CC : %v :%v\n", str8, re.MatchString(str8))
	fmt.Printf("CC : %v :%v\n", str9, re.MatchString(str9))
	fmt.Printf("CC : %v :%v\n", str10, re.MatchString(str10))
}
