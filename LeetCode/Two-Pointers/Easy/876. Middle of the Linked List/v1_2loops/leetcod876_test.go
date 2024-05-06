package main

import (
	"reflect"
	"testing"
)

func TestList_FindMiddleNode(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val: 4,
								Next: &Node{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := &List{
				head: tt.fields.head,
			}

			if got := l.FindMiddleNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.FindMiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
