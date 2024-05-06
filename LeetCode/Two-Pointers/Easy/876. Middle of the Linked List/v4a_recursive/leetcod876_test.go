package main

import (
	"reflect"
	"testing"
)

func TestList_middleNode(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
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

			if got := middleNode(l.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
