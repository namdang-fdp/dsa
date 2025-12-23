package intro_to_tree

import (
	"reflect"
	"testing"

	"github.com/namdang-fdp/dsa/tree/common"
)

func TestParentsLines(t *testing.T) {
	type args struct {
		node    *common.Node
		parrent *common.Node
	}

	// Case 1: single node
	n1 := common.NewNode(1)

	// Case 2: 1 -> {2, 3}, 3 -> {4, 5}
	//        1
	//      /   \
	//     2     3
	//          / \
	//         4   5
	root := common.NewNode(1)
	n2 := common.NewNode(2)
	n3 := common.NewNode(3)
	n4 := common.NewNode(4)
	n5 := common.NewNode(5)
	root.Children = append(root.Children, n2, n3)
	n3.Children = append(n3.Children, n4, n5)

	// Case 3: deeper chain: 10 -> 20 -> 30
	a := common.NewNode(10)
	b := common.NewNode(20)
	c := common.NewNode(30)
	a.Children = append(a.Children, b)
	b.Children = append(b.Children, c)

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single node: parent nil",
			args: args{node: n1, parrent: nil},
			want: []string{
				"1 -> NULL",
			},
		},
		{
			name: "tree: 1 has 2,3 and 3 has 4,5",
			args: args{node: root, parrent: nil},
			want: []string{
				"1 -> NULL",
				"2 -> 1",
				"3 -> 1",
				"4 -> 3",
				"5 -> 3",
			},
		},
		{
			name: "chain: 10->20->30 with explicit parent for root",
			args: args{node: a, parrent: common.NewNode(999)}, // giả lập có parent
			want: []string{
				"10 -> 999",
				"20 -> 10",
				"30 -> 20",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParentsLines(tt.args.node, tt.args.parrent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParentsLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
