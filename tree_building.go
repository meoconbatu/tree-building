package tree

import (
	"errors"
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

var (
	errNonContinuous          = errors.New("records are non-continuous")
	errRootHasParent          = errors.New("root node has parent other than itself")
	errParentGreaterThanChild = errors.New("record has parent that does not exist")
)

func Build(records []Record) (*Node, error) {
	recordLength := len(records)
	if recordLength == 0 {
		return nil, nil
	}
	nodes := make([]Node, recordLength)
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	for i, r := range records {
		if i != r.ID {
			return nil, errNonContinuous
		}
		if i == 0 {
			if r.ID != r.Parent {
				return nil, errRootHasParent
			}
		} else {
			if r.ID <= r.Parent {
				return nil, errParentGreaterThanChild
			}
			nodes[i].ID = i
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
