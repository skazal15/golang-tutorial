package main

import "testing"

func TestKonsonOrVocal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "vocal",
			args: args{
				s: "a",
			},
			want: "vocal",
		},
		{name: "konsonan",
			args: args{
				s: "b",
			},
			want: "konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KonsonOrVocal(tt.args.s); got != tt.want {
				t.Errorf("KonsonOrVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
