package tree

import (
	"errors"
	"fmt"
	"sort"
)

const testVersion = 4

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}
type Nodes []*Node
type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}
func (slice Nodes) Len() int {
	return len(slice)
}

func (slice Nodes) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice Nodes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func Build(records []Record) (*Node, error) {
	recordLength := len(records)
	if recordLength == 0 {
		return nil, nil
	}
	root := &Node{}
	todo := []*Node{root}
	n := 1
	for len(todo) > 0 {
		newTodo := []*Node(nil)
		for _, c := range todo {
			for i := len(records) - 1; i >= 0; i-- {
				r := records[i]
				if r.ID >= recordLength {
					return nil, errors.New("The ID number must be between 0 (inclusive) and the length of the record list (exclusive)")
				}
				if r.Parent == c.ID {
					switch {
					case r.ID < c.ID:
						return nil, errors.New("Higher id parent of lower id")
					case r.ID == c.ID:
						if r.ID != 0 {
							return nil, fmt.Errorf("Only root record has a parent ID that's equal to its own ID")
						}
					default:
						n++
						nn := &Node{ID: r.ID}
						newTodo = append(newTodo, nn)
						c.Children = append(c.Children, nn)
						records = remove(records, i)
					}
				}
			}
			sort.Sort(Nodes(c.Children))
		}
		todo = newTodo
	}
	if n != recordLength {
		return nil, Mismatch{}
	}
	return root, nil
}
func remove(s []Record, i int) []Record {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
