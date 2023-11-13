package dp

import (
	"fmt"
	"testing"
)

func Test_minMoney1(t *testing.T) {
	type args struct {
		arr []int
		aim int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{2, 3, 5},
				aim: 20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoney1(tt.args.arr, tt.args.aim); got != tt.want {
				t.Errorf("minMoney1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMoney2(t *testing.T) {
	type args struct {
		arr []int
		aim int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{2, 3, 5},
				aim: 20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoney2(tt.args.arr, tt.args.aim); got != tt.want {
				t.Errorf("minMoney2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_one(t *testing.T) {
	s := "1234"
	fmt.Println(s[3:])
	fmt.Println(s[4:])
}
