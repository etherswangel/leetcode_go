package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"bat"}, {"eat", "tea", "ate"}, {"tan", "nat"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			for _, g := range got {
				sort.Strings(g)
			}
			for _, w := range tt.want {
				sort.Strings(w)
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i][0] < tt.want[j][0]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
