package main

import "testing"

func TestSum(t *testing.T){
	tables := []struct{
		x int
		y int
		n int
	}{
		{1,2,3},
		{2,2,4},
		{3,3,6},
		{5,5,10},
	}

	for _, table:= range tables{
		total := Sum(table.x,table.y)
		if total!- table.n{
			t.Errorf("sum of %d + %d was incorrect, got: %d want: %d",table.x,table.y,total,table.n)
		}
	}
}