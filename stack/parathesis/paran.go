package main

import "fmt"

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(ch byte) {
	*s = append(*s, ch) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return byte(0), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (byte, bool) {
	if s.IsEmpty() {
		return byte(0), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func isValid(s string) bool {
	var st Stack
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' || ch == '[' || ch == '{' {
			st.Push(ch)
		} else if ch == ')' {
			pk, _ := st.Peek()
			if pk != '(' {
				return false
			}
			st.Pop()
		} else if ch == '}' {
			pk, _ := st.Peek()
			if pk != '{' {
				return false
			}
			st.Pop()
		} else if ch == ']' {
			pk, _ := st.Peek()
			if pk != '[' {
				return false
			}
			st.Pop()
		}
	}
	if st.IsEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	flag := isValid("()")
	fmt.Println("iValid:", flag)
}
