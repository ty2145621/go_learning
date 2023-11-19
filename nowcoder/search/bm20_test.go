package search

import (
	"testing"
)

func TestInversePairs(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				data: []int{1, 2, 3, 4, 5, 6, 7, 0},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InversePairs(tt.args.data); got != tt.want {
				t.Errorf("InversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowbit(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				x: 11,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				x: 12,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				x: 16,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowbit(tt.args.x); got != tt.want {
				t.Errorf("lowbit() = %v, want %v", got, tt.want)
			}
		})
	}
}
