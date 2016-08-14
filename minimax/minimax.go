package minimax

// Node represents a possible move
type Node struct {
	score      *int
	parent     *Node
	children   []*Node
	isOpponent bool
	data       string
}

// New returns a new minimax structure
func New() Node {
	n := Node{isOpponent: false}
	return n
}

func (node *Node) add(score *int, data string) *Node {
	childNode := Node{parent: node, score: score, data: data}

	childNode.isOpponent = !node.isOpponent
	node.children = append(node.children, &childNode)
	return &childNode
}

// AddTerminal a child node
func (node *Node) AddTerminal(score int, data string) *Node {
	return node.add(&score, data)
}

// Add normal
func (node *Node) Add(data string) *Node {
	return node.add(nil, data)
}

func (node *Node) isTerminal() bool {
	return len(node.children) == 0
}

// Evaluate minimax
func (node *Node) Evaluate() {
	eval(node)
}

func eval(node *Node) {
	for _, cn := range node.children {
		if !cn.isTerminal() {
			eval(cn)
		}

		if cn.parent.score == nil {
			cn.parent.score = cn.score
		} else if cn.isOpponent && *cn.score > *cn.parent.score {
			cn.parent.score = cn.score
		} else if !cn.isOpponent && *cn.score < *cn.parent.score {
			cn.parent.score = cn.score
		}
	}
}
