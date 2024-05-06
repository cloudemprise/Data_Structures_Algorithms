package main

import (
	"reflect"
	"testing"
)

func Test_get1stMiddleNode(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "TEST1", // Empty list
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "TEST2", // Just one node
			args: args{
				head: &Node{
					Val:  1,
					Next: nil,
				},
			},
			want: &Node{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "TEST3", // Two nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "TEST4", // Three nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &Node{
				Val: 2,
				Next: &Node{
					Val:  3,
					Next: nil,
				},
			},
		},
		{
			name: "TEST5", // Four nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		{
			name: "TEST6", // Five nodes
			args: args{
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
		{
			name: "TEST7", // Six nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val: 4,
								Next: &Node{
									Val: 5,
									Next: &Node{
										Val:  6,
										Next: nil,
									},
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
						Val: 5,
						Next: &Node{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := get1stMiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("get1stMiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

///

func Test_get2ndMiddleNode(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "TEST1", // Empty list
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "TEST2", // Just one node
			args: args{
				head: &Node{
					Val:  1,
					Next: nil,
				},
			},
			want: &Node{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "TEST3", // Two nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &Node{
				Val:  2,
				Next: nil,
			},
		},
		{
			name: "TEST4", // Three nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &Node{
				Val: 2,
				Next: &Node{
					Val:  3,
					Next: nil,
				},
			},
		},
		{
			name: "TEST5", // Four nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &Node{
				Val: 3,
				Next: &Node{
					Val:  4,
					Next: nil,
				},
			},
		},
		{
			name: "TEST6", // Five nodes
			args: args{
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
		{
			name: "TEST7", // Six nodes
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val: 4,
								Next: &Node{
									Val: 5,
									Next: &Node{
										Val:  6,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			want: &Node{
				Val: 4,
				Next: &Node{
					Val: 5,
					Next: &Node{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := get2ndMiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("get2ndMiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
