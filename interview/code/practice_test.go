package code

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *listNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *listNode
	}{
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5, 6, 7}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 6, 5, 4, 7}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5, 6}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 6, 5, 4}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 4, 5}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5}),
				k:    1,
			},
			want: build([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroupWithLast(t *testing.T) {
	type args struct {
		head *listNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *listNode
	}{
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5, 6, 7}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 6, 5, 4, 7}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5, 6}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 6, 5, 4}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: build([]int{3, 2, 1, 5, 4}),
		},
		{
			name: "",
			args: args{
				head: build([]int{1, 2, 3, 4, 5}),
				k:    1,
			},
			want: build([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroupWithLast(tt.args.head, tt.args.k); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)

			}
		})
	}
}
