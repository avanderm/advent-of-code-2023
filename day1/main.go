package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func interpret(arg string) string {
	switch arg {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return arg
	}
}

func findFirst(s string, matcher *regexp.Regexp) string {
	return matcher.FindString(s)
}

func findLast(s string, matcher *regexp.Regexp) string {
	for i := 0; i <= len(s); i++ {
		sl := s[len(s)-i:]

		if result := matcher.FindString(sl); result != "" {
			return result
		}
	}

	return ""
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	matcher, err := regexp.Compile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	check(err)

	out, err := os.Create("output")
	check(err)
	writer := bufio.NewWriter(out)

	defer out.Close()

	for scanner.Scan() {
		line := scanner.Text()
		first := findFirst(line, matcher)
		last := findLast(line, matcher)

		if first != "" && last != "" {
			first_digit := interpret(first)
			last_digit := interpret(last)

			combination := first_digit + last_digit
			d, err := strconv.ParseInt(combination, 10, 0)
			check(err)

			_, err = writer.Write([]byte(strconv.FormatInt(d, 10) + "\n"))
			check(err)

			sum += int(d)
		}
	}

	fmt.Println("sum: ", sum)
	err = writer.Flush()
	check(err)
}
