package main

func main() {

}

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
