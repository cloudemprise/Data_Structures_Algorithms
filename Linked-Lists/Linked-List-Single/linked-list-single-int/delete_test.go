package list

import (
	"testing"
)

func TestList_DeleteHead(t *testing.T) {
	type tList struct {
		head *node
	}
	tests := []struct {
		name  string
		tList tList
	}{
		// TODO: Add test cases.
		{
			name: "test",
			tList: tList{
				head: &node{
					val: 1,
					next: &node{
						val: 2,
						next: &node{
							val:  3,
							next: nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := &List{
				head: tt.tList.head,
			}

			want := &List{
				head: &node{
					val: 2,
					next: &node{
						val:  3,
						next: nil,
					},
				},
			}

			l.DeleteHead()

			if !l.isEqual(want) {
				t.Errorf("List.DeleteHead() = %v, want %v", l, want)
			}
		})
	}
}

func TestList_DeleteTail(t *testing.T) {
	type tList struct {
		head *node
	}
	tests := []struct {
		name  string
		tList tList
	}{
		// TODO: Add test cases.
		{
			name: "test",
			tList: tList{
				head: &node{
					val: 1,
					next: &node{
						val: 2,
						next: &node{
							val:  3,
							next: nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := &List{
				head: tt.tList.head,
			}

			want := &List{
				head: &node{
					val: 1,
					next: &node{
						val:  2,
						next: nil,
					},
				},
			}

			l.DeleteTail()

			if !l.isEqual(want) {
				t.Errorf("List.DeleteTail() = %v, want %v", l, want)
			}
		})
	}
}

func TestList_DeleteWithin(t *testing.T) {
	type tList struct {
		head *node
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name  string
		tList tList
		args  args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			tList: tList{
				head: &node{
					val: 1,
					next: &node{
						val: 2,
						next: &node{
							val:  3,
							next: nil,
						},
					},
				},
			},
			args: args{
				pos: 2,
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			l := &List{
				head: tt.tList.head,
			}

			want := &List{
				head: &node{
					val: 1,
					next: &node{
						val:  3,
						next: nil,
					},
				},
			}

			l.DeleteWithin(tt.args.pos)

			if !l.isEqual(want) {
				t.Errorf("List.DeleteWithin() = %v, want %v", l, want)
			}

		})
	}
}
