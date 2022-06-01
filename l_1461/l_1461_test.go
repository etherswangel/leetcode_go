package leetcode

import "testing"

func Test_hasAllCodes(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1",
			args{"00110110", 2},
			true},
		{"2",
			args{"00110", 2},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAllCodes(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("hasAllCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
