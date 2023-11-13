package dp

import (
	"testing"
)

func TestLCS1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s1: "1A2C3D4B56",
				s2: "B1D23A456A",
			},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCS1(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LCS1() = %v, want %v", got, tt.want)
			}
		})
	}
}
