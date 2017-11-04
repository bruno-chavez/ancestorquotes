package main

func SecondsMinutes() []int {
	slice := make([]int, 0)

	for i := 0;i<60;i++{
		slice = append(slice, i)
	}

	println(slice)
	return slice
}