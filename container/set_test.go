package container

import (
	"fmt"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	type args struct {
		e int
	}
	h := NewHashSet()
	tests := []struct {
		name string
		h    Set
		args args
		want bool
	}{
		{name: "v1", h: h, args: args{e: 1}, want: true},
		{name: "v2", h: h, args: args{e: 1}, want: false},
		{name: "v3", h: h, args: args{e: 2}, want: true},
		{name: "v4", h: h, args: args{e: 3}, want: true},
		{name: "v5", h: h, args: args{e: 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Add(tt.args.e); got != tt.want {
				t.Errorf("HashSet.Add() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestHashSet_Some(t *testing.T) {
	set := NewHashSet()
	set.Add(1)
	fmt.Println("set contains 1, result:", set.Contains(1))
	set.Add(2)
	fmt.Println("set contains 2, result:", set.Contains(2))
	set.Add(2)
	fmt.Println("set add 2 failed :", set.Add(2))
}
