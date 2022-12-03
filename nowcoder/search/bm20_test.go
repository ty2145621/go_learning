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
