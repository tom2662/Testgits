package applib

import (
	"math"
	"strings"
)

// GetMaxLength ...
func GetMaxLength(arr []int) float64 {
	var count float64 = 0
	var result float64 = 0

	n := len(arr)

	for i := 0; i < n; i++ {

		if arr[i] == 0 {
			count = 0
		} else {

			count++
			result = math.Max(result, count)
		}
	}

	return result
}

// Reverse ...
func Reverse(s []string) []string {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

var (
	mapParentheses = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	openParentheses  = "({["
	closeParentheses = ")}]"
)

// Element ...
type Element struct {
	value interface{}
	next  *Element
}

// Stack ...
type Stack struct {
	top  *Element
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

// Stack Pop ...
func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

// FindBalance ...
func FindBalance(parentheses string) string {

	var isValid bool
	stack := new(Stack)

	if len(parentheses) > 0 {
		isValid = true
		for _, p := range strings.Split(parentheses, "") {
			if strings.Index(openParentheses, p) != -1 {
				stack.Push(p)
				continue
			}
			if strings.Index(closeParentheses, p) != -1 {
				if stack.Len() > 0 {
					lastParentheses := stack.Pop().(string)
					if mapParentheses[lastParentheses] == p {
						continue
					}
				}
				isValid = false
				break
			}
		}
	}

	if isValid && stack.Len() == 0 {
		return "yes"
	} else {
		return "no"
	}
}
