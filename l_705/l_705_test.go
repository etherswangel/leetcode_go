package leetcode

import (
	"reflect"
	"testing"
)

func TestMyHashSet(t *testing.T) {
	type args struct {
		commands []string
		keys     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1",
			args{[]string{"MyHashSet", "add", "add", "add", "add"},
				[]int{0, 1, 2, 3, 4}},
			[]int{0, -1, -1, -1, -1},
		},
		{"2",
			args{[]string{"MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"},
				[]int{0, 1, 2, 1, 3, 2, 2, 2, 2}},
			[]int{0, -1, -1, 1, 0, -1, 1, -1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			hash := Constructor()
			for idx, command := range tt.args.commands {
				if command == "add" {
					hash.Add(tt.args.keys[idx])
					got = append(got, -1)
				} else if command == "remove" {
					hash.Remove(tt.args.keys[idx])
					got = append(got, -1)
				} else {
					res := hash.Contains(tt.args.keys[idx])
					if res == true {
						got = append(got, 1)
					} else {
						got = append(got, 0)
					}
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name = %v, got = %v, want %v", tt.name, got, tt.want)
			} else {
				t.Logf("Passed, Name = %v, got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
