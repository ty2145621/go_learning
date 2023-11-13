package recursion

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				num: []int{2, 2, -1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPermutation(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		{name: "", args: args{num: []int{1, 2, 3, 4}}, want: true, want1: []int{1, 2, 4, 3}},
		{name: "", args: args{num: []int{4, 2, 1, 3}}, want: true, want1: []int{4, 2, 3, 1}},
		{name: "", args: args{num: []int{3, 4, 1, 2}}, want: true, want1: []int{3, 4, 2, 1}},
		{name: "", args: args{num: []int{3, 4, 2, 1}}, want: true, want1: []int{4, 1, 2, 3}},
		{name: "", args: args{num: []int{3, 2, 4, 1}}, want: true, want1: []int{3, 4, 1, 2}},
		{name: "", args: args{num: []int{3, 2, 1, 4}}, want: true, want1: []int{3, 2, 4, 1}},
		{name: "", args: args{num: []int{4, 3, 2, 1}}, want: false, want1: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nextPermutation(tt.args.num)
			if got != tt.want {
				t.Errorf("nextPermutation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("nextPermutation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_permuteUnique1(t *testing.T) {
	type args struct {
		num    []int
		result [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				num:    []int{1, 2, 3},
				result: make([][]int, 0),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique1(tt.args.num, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique1() = %v, want %v", got, tt.want)
			}
		})
	}
}
