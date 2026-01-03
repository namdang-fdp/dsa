package intro_to_tree

import (
	"fmt"
	"strings"

	"github.com/namdang-fdp/dsa/tree/common"
)

func AddChild(children, parent *common.Node) {
	parent.Children = append(parent.Children, children)
}

func ParentsLines(node, parrent *common.Node) []string {
	var out []string
	var dfs func(node, parrent *common.Node)
	dfs = func(node, parrent *common.Node) {
		if parrent == nil {
			out = append(out, fmt.Sprintf("%d -> NULL", node.Data))
		} else {
			out = append(out, fmt.Sprintf("%d -> %d", node.Data, parrent.Data))
		}
		for _, ch := range node.Children {
			dfs(ch, node)
		}
	}
	dfs(node, parrent)
	return out
}

// ChildrenLines returns: "node -> child1 child2 ..."
func ChildrenLines(node *common.Node) []string {
	var out []string
	var dfs func(n *common.Node)
	dfs = func(n *common.Node) {
		var b strings.Builder
		b.WriteString(fmt.Sprintf("%d -> ", n.Data))
		for _, ch := range n.Children {
			b.WriteString(fmt.Sprintf("%d ", ch.Data))
		}
		out = append(out, strings.TrimSpace(b.String()))
		for _, ch := range n.Children {
			dfs(ch)
		}
	}
	dfs(node)
	return out
}

func LeafValues(node *common.Node) []int {
	var out []int
	var dfs func(n *common.Node)
	dfs = func(n *common.Node) {
		if len(n.Children) == 0 {
			out = append(out, n.Data)
			return
		}
		for _, ch := range n.Children {
			dfs(ch)
		}
	}
	dfs(node)
	return out
}

// DegreesLines returns: "node -> degree" (children + parent edge if any)
func DegreesLines(node, parent *common.Node) []string {
	var out []string
	var dfs func(n, p *common.Node)
	dfs = func(n, p *common.Node) {
		deg := len(n.Children)
		if p != nil {
			deg++
		}
		out = append(out, fmt.Sprintf("%d -> %d", n.Data, deg))
		for _, ch := range n.Children {
			dfs(ch, n)
		}
	}
	dfs(node, parent)
	return out
}
