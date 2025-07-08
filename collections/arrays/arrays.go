package main

import "fmt"

func refactored_array() {
	var MoreDeploymentOptions = [...]string{"AWS","GCP"}
	
	// Compiler Counts the Size

	for index, option := range MoreDeploymentOptions {
		fmt.Println(index, option)
	}
}

func main() {
	var DeploymentOptions = [3]string{"AWS", "Azure", "GCP"}

	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}

	fmt.Println("-------------")

	refactored_array()
}


