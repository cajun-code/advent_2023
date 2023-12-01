package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ProcessFileWithRegex(file string) int64 {
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
		val := stringToDigitRegex(line)
		fmt.Printf("Numbers found: %s\t", val)
		format := "%s%s"
		rune_count := len(val)
		value := ""
		switch rune_count {
		case 0:
			continue
		case 1:
			value = fmt.Sprintf(format, val[0], val[0])
		default:
			value = fmt.Sprintf(format, val[0], val[rune_count-1])
		}
		fmt.Printf("Processed val: %s\n", value)
		convert, err := strconv.ParseInt(value, 10, 32)
		Check(err)
		total += convert

	}
	fmt.Printf("\nTotal: %d\n\n", total)
	return total
}

func stringToDigitRegex(digit string) []string {
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	exp := fmt.Sprintf("%s", strings.Join(append(digits, words...), "|"))
	rx := regexp.MustCompile(exp)
	values := rx.FindAllString(digit, -1)
	//fmt.Println(values)
	for i, num := range values {
		if contains(words, num) {
			pos := indexOf(num, words)
			values[i] = digits[pos]
		}
	}
	return values
}
