package code

import (
	"testing"
)

func TestPrintAB(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				n: 6,
			},
			want: "12a4ba",
		},
		{
			name: "",
			args: args{
				n: ,
			},
			want: "12a4ba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintAB(tt.args.n); got != tt.want {
				t.Errorf("PrintAB() = %v, want %v", got, tt.want)
			}
		})
	}
}
