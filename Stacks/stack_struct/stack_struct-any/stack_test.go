package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type tStack struct {
		items []any
	}
	type args struct {
		item any
	}
	tests := []struct {
		name   string
		tStack tStack
		args   args
	}{
		// test cases:
		{
			name: "test",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
			args: args{'d'},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			s := &Stack{
				items: tt.tStack.items,
			}

			want := &Stack{
				items: []any{'a', 'b', 'c', 'd'},
			}

			if s.Push(tt.args.item); !reflect.DeepEqual(s, want) {
				t.Errorf("Push() = %c, want %c", s, want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type tStack struct {
		items []any
	}
	tests := []struct {
		name   string
		tStack tStack
		want   any
	}{
		// test cases:
		{
			name: "test",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
			want: 'c',
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			s := &Stack{
				items: tt.tStack.items,
			}

			if got := s.Pop(); got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type tStack struct {
		items []any
	}
	tests := []struct {
		name   string
		tStack tStack
		want   any
	}{
		// test cases:
		{
			name: "test",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items: tt.tStack.items,
			}
			if got := s.Peek(); got != tt.want {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type tStack struct {
		items []any
	}
	tests := []struct {
		name   string
		tStack tStack
		want   bool
	}{
		// test cases:
		{
			name: "test1",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
			want: false,
		},
		{
			name: "test2",
			tStack: tStack{
				items: []any{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items: tt.tStack.items,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	type tStack struct {
		items []any
	}
	tests := []struct {
		name   string
		tStack tStack
		want   int
	}{
		// test cases:
		{
			name: "test1",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
			want: 3,
		},
		{
			name: "test2",
			tStack: tStack{
				items: []any{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items: tt.tStack.items,
			}
			if got := s.Length(); got != tt.want {
				t.Errorf("Stack.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Reset(t *testing.T) {
	type tStack struct {
		items []any
	}
	tests := []struct {
		name   string
		tStack tStack
	}{
		// test cases:
		{
			name: "test",
			tStack: tStack{
				items: []any{'a', 'b', 'c'},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			s := &Stack{
				items: tt.tStack.items,
			}

			want := &Stack{}

			if s.Reset(); !reflect.DeepEqual(s, want) {
				t.Errorf("Reset() = %c, want %c", s, want)
			}
		})
	}
}
