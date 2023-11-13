package real

import (
	"reflect"
	"testing"
)

func TestLFU2(t *testing.T) {
	type args struct {
		operators [][]int
		k         int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				operators: [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {1, 2, 4}, {1, 3, 5}, {2, 2}, {1, 4, 4}, {2, 1}},
				k:         3,
			},
			want: []int{4, -1},
		},
		{
			name: "",
			args: args{
				operators: [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 3}, {1, 4, 4}, {2, 4}, {2, 3}, {2, 2}, {2, 1}, {1, 5, 5}, {2, 4}},
				k:         4,
			},
			want: []int{4, 3, 2, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LFU2(tt.args.operators, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LFU2() = %v, want %v", got, tt.want)
			}
		})
	}
}
