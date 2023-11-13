package stack

import (
	"fmt"
	"sort"
	"testing"
)

func Test_findKth(t *testing.T) {
	type args struct {
		a []int
		n int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: []int{3, 4, 1, 2, 5}, n: 5, K: 1}, want: 5},
		{name: "2", args: args{a: []int{3, 4, 1, 2, 5}, n: 5, K: 2}, want: 4},
		{name: "3", args: args{a: []int{3, 4, 1, 2, 5}, n: 5, K: 3}, want: 3},
		{name: "4", args: args{a: []int{3, 4, 1, 2, 5}, n: 5, K: 4}, want: 2},
		{name: "5", args: args{a: []int{3, 4, 1, 2, 5}, n: 5, K: 5}, want: 1},
		{
			name: "",
			args: args{
				a: []int{1332802, 1177178, 1514891, 871248, 753214, 123866, 1615405, 328656, 1540395, 968891,
					1884022, 252932, 1034406, 1455178, 821713, 486232, 860175, 1896237, 852300, 566715, 1285209,
					1845742, 883142, 259266, 520911, 1844960, 218188, 1528217, 332380, 261485, 1111670, 16920,
					1249664, 1199799, 1959818, 1546744, 1904944, 51047, 1176397, 190970, 48715, 349690, 673887,
					1648782, 1010556, 1165786, 937247, 986578, 798663},
				n: 49,
				K: 24,
			},
			want: 986578,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKth(tt.args.a, tt.args.n, tt.args.K); got != tt.want {
				sort.Reverse(sort.IntSlice(tt.args.a))
				fmt.Println(tt.args.a)
				t.Errorf("findKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
