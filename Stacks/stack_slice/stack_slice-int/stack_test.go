package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		tStack *Stack
		args   args
	}{
		// test cases:
		{
			name:   "test",
			tStack: &Stack{1, 2, 3},
			args:   args{4},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			want := &Stack{1, 2, 3, 4}

			if tt.tStack.Push(tt.args.item); !reflect.DeepEqual(tt.tStack, want) {
				t.Errorf("Push() = %c, want %c", *tt.tStack, want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name   string
		tStack *Stack
		want   int
	}{
		// test cases:
		{
			name:   "test",
			tStack: &Stack{1, 2, 3},
			want:   3,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := tt.tStack.Pop(); got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name   string
		tStack *Stack
		want   int
	}{
		// test cases:
		{
			name:   "test",
			tStack: &Stack{1, 2, 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.tStack.Peek(); got != tt.want {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {

	tests := []struct {
		name   string
		tStack *Stack
		want   bool
	}{
		// test cases:
		{
			name:   "test",
			tStack: &Stack{1, 2, 3},
			want:   false,
		},
		{
			name:   "test",
			tStack: &Stack{},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.tStack.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	tests := []struct {
		name   string
		tStack *Stack
		want   int
	}{
		// test cases:
		{
			name:   "test1",
			tStack: &Stack{1, 2, 3},
			want:   3,
		},
		{
			name:   "test2",
			tStack: &Stack{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.tStack.Length(); got != tt.want {
				t.Errorf("Stack.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Reset(t *testing.T) {

	tests := []struct {
		name   string
		tStack *Stack
	}{
		// test cases:
		{
			name:   "test",
			tStack: &Stack{1, 2, 3},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			want := &Stack{}

			if tt.tStack.Reset(); !reflect.DeepEqual(tt.tStack, want) {
				t.Errorf("Reset() = %c, want %c", *tt.tStack, want)
			}
		})
	}
}
