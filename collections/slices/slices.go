package main

import "fmt"

func main() {
	
	languages := [5]string{"java", "python", "go", "SQL", "c",}

	// Sclices are dynamic Arrays
	// Define Slices
	// Slice is view of original Array
	// If array full slice is 
	
	firstThree := languages[0:3]
	firstThree[2] = "c#"
	// firstThree points to languages array
	fmt.Println(languages)


	lastTwoLangs := languages[3:5]
	lastTwoLangs = append(lastTwoLangs, "abc")

	fmt.Println(lastTwoLangs)
	// appending to lastTwoLangs Slice didn't affect original langs array
	fmt.Println(languages)
	

	boundedArray := [4]int{1, 2, 3, 4}
	newSlice := make([]int, len(boundedArray), 10)

	// Only Copies One Slice to another slice
	copy(newSlice, boundedArray[:])
	
	newSlice = append(newSlice, 5)
	fmt.Println(newSlice)

	

}
