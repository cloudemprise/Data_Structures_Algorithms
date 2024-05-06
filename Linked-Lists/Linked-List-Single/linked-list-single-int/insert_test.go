package list

import "testing"

// isEqual checks if two linked list are equal.
func (l *List) isEqual(other *List) bool {
	current1 := l.head
	current2 := other.head

	for current1 != nil && current2 != nil {
		if current1.val != current2.val {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}

	// If one reached the end and the other didn't, they are not equal
	if current1 != current2 {
		return false
	}

	// If we reached the ends together, they are equal
	return true
}

func TestList_InsertHead(t *testing.T) {
	type tList struct {
		head *node
	}
	type args struct {
		val int
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
						val:  2,
						next: nil,
					},
				},
			},
			args: args{
				val: 3,
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
					val: tt.args.val,
					next: &node{
						val: 1,
						next: &node{
							val:  2,
							next: nil,
						},
					},
				},
			}

			l.InsertHead(tt.args.val)

			if !l.isEqual(want) {
				t.Errorf("List.InsertHead() = %v, want %v", l, want)
			}

		})
	}
}

func TestList_InsertTail(t *testing.T) {
	type tList struct {
		head *node
	}
	type args struct {
		val int
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
						val:  2,
						next: nil,
					},
				},
			},
			args: args{
				val: 3,
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
						val: 2,
						next: &node{
							val:  tt.args.val,
							next: nil,
						},
					},
				},
			}

			l.InsertTail(tt.args.val)

			if !l.isEqual(want) {
				t.Errorf("List.InsertTail() = %v, want %v", l, want)
			}

		})
	}
}

func TestList_InsertWithin(t *testing.T) {
	type tList struct {
		head *node
	}
	type args struct {
		val int
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
						val:  2,
						next: nil,
					},
				},
			},
			args: args{
				val: 3,
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
						val: tt.args.val,
						next: &node{
							val:  2,
							next: nil,
						},
					},
				},
			}

			l.InsertWithin(tt.args.val, tt.args.pos)

			if !l.isEqual(want) {
				t.Errorf("List.InsertWithin() = %v, want %v", l, want)
			}

		})
	}
}
