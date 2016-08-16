package minimax

import (
	"fmt"
	"strconv"
)

// Node represents a possible score
type Node struct {
	Score      *int
	parent     *Node
	children   []*Node
	isOpponent bool
	Data       interface{}
}

// New returns a new minimax structure
func New() Node {
	n := Node{isOpponent: false}
	return n
}

// GetBestChildNode returns the first child node with the matching score
func (node *Node) GetBestChildNode() *Node {
	for _, cn := range node.children {
		if cn.Score == node.Score {
			return cn
		}
	}

	return nil
}

// Evaluate minimax
func (node *Node) Evaluate() {
	for _, cn := range node.children {
		if !cn.isTerminal() {
			cn.Evaluate()
		}

		if cn.parent.Score == nil {
			cn.parent.Score = cn.Score
		} else if cn.isOpponent && *cn.Score > *cn.parent.Score {
			cn.parent.Score = cn.Score
		} else if !cn.isOpponent && *cn.Score < *cn.parent.Score {
			cn.parent.Score = cn.Score
		}
	}
}

// Print the node
func (node *Node) Print(level int) {
	var padding = ""
	for j := 0; j < level; j++ {
		padding += " "
	}

	var s = ""
	if node.Score != nil {
		s = strconv.Itoa(*node.Score)
	}

	fmt.Println(padding, node.isOpponent, node.Data, "["+s+"]")

	for _, cn := range node.children {
		level += 2
		cn.Print(level)
		level -= 2
	}

}

// AddTerminal a child node
func (node *Node) AddTerminal(score int, data int) *Node {
	return node.add(&score, data)
}

// Add normal
func (node *Node) Add(data int) *Node {
	return node.add(nil, data)
}

func (node *Node) add(score *int, data int) *Node {
	childNode := Node{parent: node, Score: score, Data: data}

	childNode.isOpponent = !node.isOpponent
	node.children = append(node.children, &childNode)
	return &childNode
}

func (node *Node) isTerminal() bool {
	return len(node.children) == 0
}
