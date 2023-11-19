package practice

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func initArr() []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 0, 20)
	for i := 0; i <= 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				arr: initArr(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args.arr)
			quickSort2(tt.args.arr)
			t.Logf("\n%+v", tt.args.arr)
		})
	}
}

func Test_reverseList(t *testing.T) {
	head := &listNode{}
	head.build([]int{1, 2, 3, 4, 5})

	type args struct {
		head *listNode
	}
	tests := []struct {
		name string
		args args
		want *listNode
	}{
		{
			name: "",
			args: args{
				head: head,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args.head)
			got := reverseList(tt.args.head)
			t.Logf("\n%+v", got)
		})
	}
}

func Test_reverseKList(t *testing.T) {
	head := &listNode{}
	head.build([]int{1, 2, 3, 4, 5})
	head1 := &listNode{}
	head1.build([]int{1, 2, 3, 4, 5})
	head2 := &listNode{}
	head2.build([]int{1, 2, 3, 4, 5})

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
				head: head,
				k:    0,
			},
		},
		{
			name: "",
			args: args{
				head: head,
				k:    1,
			},
		},
		{
			name: "",
			args: args{
				head: head1,
				k:    2,
			},
		},
		{
			name: "",
			args: args{
				head: head2,
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args.head)
			got := reverseKList(tt.args.head, tt.args.k)
			t.Logf("%+v", got)
		})
	}
}

func Test_reverseKList2(t *testing.T) {
	head := &listNode{}
	head.build([]int{1, 2, 3, 4, 5})
	head1 := &listNode{}
	head1.build([]int{1, 2, 3, 4, 5})
	head2 := &listNode{}
	head2.build([]int{1, 2, 3, 4, 5})
	head3 := &listNode{}
	head3.build([]int{1, 2, 3, 4, 5})
	head4 := &listNode{}
	head4.build([]int{1, 2, 3, 4, 5})

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
				head: head,
				k:    0,
			},
		},
		{
			name: "",
			args: args{
				head: head,
				k:    1,
			},
		},
		{
			name: "",
			args: args{
				head: head1,
				k:    2,
			},
		},
		{
			name: "",
			args: args{
				head: head2,
				k:    3,
			},
		},
		{
			name: "",
			args: args{
				head: head3,
				k:    4,
			},
		},
		{
			name: "",
			args: args{
				head: head4,
				k:    5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args.head)
			got := reverseKList2(tt.args.head, tt.args.k)
			t.Logf("%+v", got)
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 5,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	head := &listNode{}
	head.build([]int{1, 3, 5, 7, 9})
	head1 := &listNode{}
	head1.build([]int{2, 5, 8, 11, 14})
	head2 := &listNode{}
	head2.build([]int{3, 4, 6, 8, 10})

	headList := []*listNode{head, head1, head2}


	type args struct {
		l []*listNode
	}
	tests := []struct {
		name string
		args args
		want *listNode
	}{
		{
			name: "",
			args: args{
				l: headList,
			},
			want: &listNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
