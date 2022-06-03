package main

import (
	"errors"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	name          string
	worth, weight int
}

type bag struct {
	bagWeight, currItemsWeight, maxItemsWeight, totalWeight, totalWorth int
	items                                                               []item
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readItems(path string) []item {

	dat, err := ioutil.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")

	itemList := make([]item, 0)

	for i, v := range lines {
		if i == 0 {
			continue // skip the headers on the first line
		}
		s := strings.Split(v, ",")
		newItemWorth, _ := strconv.Atoi(s[1])
		weigth := strings.TrimSuffix(s[2], "\r")
		newItemWeight, _ := strconv.Atoi(weigth)
		newItem := item{name: s[0], worth: newItemWorth, weight: newItemWeight}
		itemList = append(itemList, newItem)
	}
	return itemList
}

func (b *bag) addItem(i item) error {
	if b.currItemsWeight+i.weight <= b.maxItemsWeight {
		b.currItemsWeight += i.weight
		b.items = append(b.items, i)
		println(i.name)
		return nil
	}
	return errors.New("could not fit item ")
}

func greedy(is []item, b bag) {
	println("output from greedy algorithm")
	sort.Slice(is, func(i, j int) bool {
		return is[i].worth > is[j].worth
	})

	for i := range is {
		b.addItem(is[i])
	}

	b.totalWeight = b.bagWeight + b.currItemsWeight

	for _, v := range b.items {
		b.totalWorth += v.worth
	}

}
