package interview

import (
	"reflect"
	"testing"
)

func Test_getNext(t *testing.T) {
	type args struct {
		n     int
		limit int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "", args: args{n: 14, limit: 123},
			want: 15, want1: true,
		},
		{
			name: "", args: args{n: 10, limit: 123},
			want: 100, want1: true,
		},
		{
			name: "", args: args{n: 101, limit: 123},
			want: 102, want1: true,
		},
		{
			name: "", args: args{n: 109, limit: 123},
			want: 11, want1: true,
		},
		{
			name: "", args: args{n: 99, limit: 123},
			want: -1, want1: false,
		},
		{
			name: "", args: args{n: 123, limit: 123},
			want: 13, want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getNext(tt.args.n, tt.args.limit)
			if got != tt.want {
				t.Errorf("getNext() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getNext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_printOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				n: 123,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				n: 1003,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
