package interview

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				num: []int{3, 2, 1, 6, 5},
			},
		},
		{name: "", args: args{num: []int{}}},
		{name: "", args: args{num: []int{1}}},
		{name: "", args: args{num: []int{2, 1, 2, 6, 5, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.num)
			t.Logf("%+v", tt.args.num)
		})
	}
}

func Test_findAll(t *testing.T) {
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
				num: []int{1, 2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAll(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
