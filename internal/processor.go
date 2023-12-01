package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ProcessFile(file string) int64 {
	var total int64
	data, err := os.ReadFile(file)
	Check(err)
	dat := string(data)
	lines := strings.Split(dat, "\n")
	for no, line := range lines {
		// Process letter of line to words
		// read first letter of string into temp
		// if digit then reset temp string
		// check if temp is digit word
		// replace digit word with numerical value
		// and reset temp string
		fmt.Printf("Line: %d %s\t", no+1, line)
		var temp string
		val := strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) {
				temp = ""
				return r
			}
			temp += string(r)
			rv := stringToDigit(temp)
			if rv > -1 {
				temp = ""
				return rv
			}
			return -1
		}, line)
		temp = ""
		fmt.Printf("Numbers found: %s\t", val)
		rune_count := utf8.RuneCountInString(val)
		runes := []rune(val)
		switch rune_count {
		case 0:
			continue
		case 1:
			val = fmt.Sprint(val, val)
		case 2:
			break
		default:
			val = fmt.Sprint(string(runes[0:1]), string(runes[rune_count-1:]))
		}
		fmt.Printf("Processed val: %s\n", val)
		value, err := strconv.ParseInt(val, 10, 32)
		Check(err)
		total += value

	}
	fmt.Printf("\nTotal: %d\n\n", total)
	return total
}

func stringToDigit(digit string) rune {
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "xxxxxx"}
	var word string
	var index int
	for index, word = range digits {
		if strings.Contains(digit, word) {
			break
		}
		if index == 10 {
			word = "notfound"
			index = 0
		}
	}
	switch strings.ToLower(word) {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	//case "zero":
	//	return '0'

	default:
		return -1
	}
}
