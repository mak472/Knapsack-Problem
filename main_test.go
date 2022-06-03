package main

import (
	"testing"
)

func Benchmark_greedy(b *testing.B) {
	itemList := readItems("objects.csv")
	for i := 0; i < b.N; i++ {
		minaal := bag{bagWeight: 1415, currItemsWeight: 0, maxItemsWeight: 5585}
		greedy(itemList, minaal)
	}
}

func Benchmark_dynamic(b *testing.B) {
	itemList := readItems("objects.csv")
	for i := 0; i < b.N; i++ {
		minaal := bag{bagWeight: 1415, currItemsWeight: 0, maxItemsWeight: 5585}
		dynamic(itemList, &minaal)
	}
}
