package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Stack interface {
	push(string) error
	pop() string
	empty() bool
}

type stack struct {
	base []brace
}

type brace struct {
	symbol   string
	position int
}

func (s *stack) push(value brace) error {
	s.base = append(s.base, value)
	return nil
}

func (s *stack) pop() brace {
	n := len(s.base) - 1 // Верхний элемент
	// fmt.Print(s[n])
	elem := s.base[n]
	// s.base[n] = "" // Удаляем элемент (записываем нулевое значение)
	s.base = s.base[:n]
	return elem
}

func (s *stack) empty() bool {
	return len(s.base) == 0
}

func main() {
	var input string

	fmt.Scan(&input)

	fmt.Println(isBalanced(input))
}

func isBalanced(str string) string {
	var openBraces = stack{
		base: make([]brace, 0),
	}

	for k, v := range str {
		v := string(v)

		if ok := regexp.MustCompile(`[(){}\[\]]`).Match([]byte(v)); !ok {
			continue
		}

		if v == "(" || v == "[" || v == "{" {
			openBraces.push(brace{symbol: string(v), position: k + 1})
		} else {
			if openBraces.empty() {
				return strconv.Itoa(k + 1)
			}

			openBrace := openBraces.pop()
			if (openBrace.symbol == "[" && v != "]") || (openBrace.symbol == "(" && v != ")") || (openBrace.symbol == "{" && v != "}") {
				return strconv.Itoa(k + 1)
			}
		}
	}

	if !openBraces.empty() {
		return strconv.Itoa(openBraces.pop().position)
	}

	return "Success"
}
