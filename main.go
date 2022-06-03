package main

func main() {

	minaal := bag{bagWeight: 130, currItemsWeight: 0, maxItemsWeight: 200}
	itemList := readItems("objects.csv")

	greedy(itemList, minaal)
	dynamic(itemList, &minaal)

	//	print(bag)
	// r := Router()
	// fmt.Println("Starting server on the port 8080...")
	// log.Fatal(http.ListenAndServe(":8080", r))
}
