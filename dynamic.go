package main

import "math"

func checkItem(b *bag, i int, w int, is []item, matrix [][]int) {
	if i <= 0 || w <= 0 {
		return
	}

	pick := matrix[i][w]
	if pick != matrix[i-1][w] {
		b.addItem(is[i-1])
		checkItem(b, i-1, w-is[i-1].weight, is, matrix)
	} else {
		checkItem(b, i-1, w, is, matrix)
	}
}

func dynamic(is []item, b *bag) *bag {
	println("output from dynamic programming")
	numItems := len(is)          // number of items in knapsack
	capacity := b.maxItemsWeight // capacity of knapsack

	// create the empty matrix
	matrix := make([][]int, numItems+1) // rows representing items
	for i := range matrix {
		matrix[i] = make([]int, capacity+1) // columns representing grams of weight
	}

	// loop through table rows
	for i := 1; i <= numItems; i++ {
		// loop through table columns
		for w := 1; w <= capacity; w++ {

			// if weight of item matching this index can fit at the current capacity column...
			if is[i-1].weight <= w {
				// worth of this subset without this item
				valueOne := float64(matrix[i-1][w])
				// worth of this subset without the previous item, and this item instead
				valueTwo := float64(is[i-1].worth + matrix[i-1][w-is[i-1].weight])
				// take maximum of either valueOne or valueTwo
				matrix[i][w] = int(math.Max(valueOne, valueTwo))
				// if the new worth is not more, carry over the previous worth
			} else {
				matrix[i][w] = matrix[i-1][w]
			}
		}
	}

	checkItem(b, numItems, capacity, is, matrix)

	// add other statistics to the bag
	b.totalWorth = matrix[numItems][capacity]
	b.totalWeight = b.bagWeight + b.currItemsWeight

	return b
}
