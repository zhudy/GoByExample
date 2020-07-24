package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int{
	if a < b {
		return a
	}
	return b
}
func TestIntMinBasic(t *testing.T){
	ans := IntMin(2, -2)
	if ans != -1{
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
func TestIntMinTableDriven(t *testing.T){
	var tests = []struct{
		a, b, want int
	}{
		{0,1,0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests{
		testname := fmt.Sprintf("%d,%d", tt.a,tt.b)
		t.Run(testname, func(t *testing.T){
			ans := IntMin(tt.a, tt.b)
			if tt.want != ans{
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
