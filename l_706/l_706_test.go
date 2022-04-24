package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMyHashSet(t *testing.T) {
	type args struct {
		commands []string
		keys     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1",
			args{[]string{"MyHashMap", "put", "put", "get", "get", "put", "get", "remove"},
				[][]int{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}}},
			[]int{-1, -1, 1, -1, -1, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			hash := Constructor()
			for idx, command := range tt.args.commands {
				if command == "put" {
					fmt.Println("put")
					hash.Put(tt.args.keys[idx][0], tt.args.keys[idx][1])
					got = append(got, -1)
				} else if command == "remove" {
					fmt.Println("rem")
					hash.Remove(tt.args.keys[idx][0])
					got = append(got, -1)
				} else if command == "get" {
					fmt.Println("get")
					res := hash.Get(tt.args.keys[idx][0])
					got = append(got, res)
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
