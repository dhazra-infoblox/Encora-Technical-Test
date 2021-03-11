package main

import (
	"encoding/json"
	"log"
)

// STACK data structure implementation open

//Stack data structure
type Stack struct {
	items     []string         // stack to parse through the input string
	nodeExist map[string]*node // keep track of already visited node by its name field
}

//Push //
func (stack *Stack) Push(item string) {
	stack.items = append(stack.items, item)
}

//Pop //
func (stack *Stack) Pop() string {
	if len(stack.items) == 0 {
		return ""
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

//IsEmpty //
func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

//Peek //
func (stack *Stack) Peek() string {
	if len(stack.items) == 0 {
		return ""
	}

	return stack.items[len(stack.items)-1]
}

//Dump //
func (stack *Stack) Dump() []string {
	var copiedStack = make([]string, len(stack.items))
	copy(copiedStack, stack.items)

	return copiedStack
}

// STACK data structure implementation close

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

func parse(v string) (*node, error) {
	root := &node{Name: ""}
	stack := Stack{nodeExist: map[string]*node{}}
	strElm := "" //string builder to create the node element
	for _, char := range v {
		if string(char) == "," {
			// string created before ',' is found should be pushed
			if strElm != "" {
				stack.Push(strElm)
				strElm = ""
			}
			//do nothing
		} else if string(char) == "]" {
			// string created before ']' is found should be pushed
			if strElm != "" {
				stack.Push(strElm)
				strElm = ""
			}
			parent := &node{Name: ""}                     // Create empty parent
			for !stack.IsEmpty() && stack.Peek() != "[" { // Keep poping elements till '[' is found
				elm := stack.Pop()
				ok, exist := stack.nodeExist[elm] // is the node already created and stored temporarily ?
				if !exist {                       // No, create a new Node and append it to the empty parent children list
					newNode := &node{Name: elm}
					//stack.nodeExist[elm] =
					parent.Children = append(parent.Children, newNode)
				} else { // Yes, fetch the node with its children and append the node to the empty parent children list
					parent.Children = append(parent.Children, ok)
				}
			}
			stack.Pop()           // Pop '['
			if !stack.IsEmpty() { // if stack not empty, the top element must be the actual Parent
				elm := stack.Peek() // Do not pop the element as it will be reused in the next phase
				parent.Name = elm
				stack.nodeExist[elm] = parent // store this parent with the elm as a key
			} else {
				root = parent // else parent should be the result
			}
		} else if string(char) == "[" {
			// string created before '[' is found should be pushed
			if strElm != "" {
				stack.Push(strElm)
				strElm = ""
			}
			stack.Push(string(char)) // push '['
		} else {
			//make the element before its pushed to stack
			strElm = strElm + string(char)
		}
	}
	return root, nil
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func main() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
