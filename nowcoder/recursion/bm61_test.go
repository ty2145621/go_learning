package recursion

import (
	"testing"
)

func Test_solve61(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				matrix: [][]int{
					{9, 8, 7},
					{6, 5, 4},
					{3, 2, 1},
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{6, 5, 4},
					{9, 8, 7},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve61(tt.args.matrix); got != tt.want {
				t.Errorf("solve61() = %v, want %v", got, tt.want)
			}
		})
	}
}
