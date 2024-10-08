package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

func main() {
	//input := []int{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6}
	input := []int{-1, 0, 1}
	fmt.Println(minimumGroups(input))
	// Block of code to be added to the main function
	// Code to be added to the main function
	// Code to be added to the main functione to create cloud storage bucket

}

// Function to create a Cloud Storage bucket
func createBucket(bucketName string) error {
	// Create a new Cloud Storage client
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create Cloud Storage client: %w", err)
	}
	defer client.Close()

	// Create the bucket
	ctx := context.Background()
	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, nil); err != nil {
		return fmt.Errorf("failed to create bucket %q: %w", bucketName, err)
	}

	fmt.Printf("Bucket %q created successfully.\n", bucketName)
	return nil
}

// Function to calculate the minimum number of groups
func minimumGroups(predators []int) int {
	// construct predatorMap
	n := len(predators)
	predatorMap := make([][]int, n)
	for i := 0; i < n; i++ {
		if predators[i] == -1 {
			continue
		}
		predatorMap[predators[i]] = append(predatorMap[predators[i]], i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if predators[i] == -1 {
			ans = max(ans, maxDepth(predatorMap, i))
		}
	}

	return ans
}

func maxDepth(predatorMap [][]int, p int) int {
	depth := 1

	for i := 0; i < len(predatorMap[p]); i++ {
		depth = max(depth, 1+maxDepth(predatorMap, predatorMap[p][i]))
	}

	return depth
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
{-1,8,6,0,7,3,8,9,-1,6}
                     ^
array pemakan
0 -> 3
1 ->
2
3 -> 5
4
5
6 -> 2,9
7 -> 4
8 -> 1,6
9 -> 7


*/
