package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByUserAge []user

func (a ByUserAge) Len() int {
	return len(a)
}

func (a ByUserAge) Less(b int, c int) bool {
	return a[b].Age < a[c].Age
}

func (a ByUserAge) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByUseLastName []user

func (a ByUseLastName) Len() int {
	return len(a)
}

func (a ByUseLastName) Less(b int, c int) bool {
	return a[b].Last < a[c].Last
}

func (a ByUseLastName) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here
	fmt.Println("Before sort")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("Sorted by age")
	sort.Sort(ByUserAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	// your code goes here
	fmt.Println("Sorted by LastName")
	sort.Sort(ByUseLastName(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}
