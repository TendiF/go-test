package HouseRobber
import (
	"fmt"
)
func SafeRob(h []int) int{
	total := 0
	for i,v := range h{
		if i % 2 == 0 {
			total += v
		}
	}
	return total
}

func RunRob(){
	total := SafeRob([]int{1,2,3,1})
	fmt.Println("Total Rob", total)
}