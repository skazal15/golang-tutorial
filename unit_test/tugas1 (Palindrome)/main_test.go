package main

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "palindrome",
			args: args{
				r: "madam",
			},
			want: true,
		},
		{
			name: "non-palindrome",
			args: args{
				r: "said",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.r); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
