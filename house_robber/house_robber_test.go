package HouseRobber

import (
	"testing"
)


func TestSafeRob(t *testing.T){
	total := SafeRob([]int{1,2,3,1})
	if total != 4 {
		t.Error("Error need to be:", 4)
	}
}