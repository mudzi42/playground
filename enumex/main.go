package main

import "fmt"

type Weekday int

// const (
// 	Sunday    Weekday = 0
// 	Monday    Weekday = 1
// 	Tuesday   Weekday = 2
// 	Wednesday Weekday = 3
// 	Thursday  Weekday = 4
// 	Friday    Weekday = 5
// 	Saturday  Weekday = 6
// )

// a different way to define the const
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	if day < Sunday || day > Saturday {
		return "Unkown day"
	}

	return names[day]
}

func (day Weekday) Weekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

type Timezone int

// const (
// 	// iota: 0, EST: -5
// 	EST Timezone = -(5 + iota)

// 	// iota: 1, CST: -6 
// 	CST

// 	// iota: 2, MST: -7
// 	MST

// 	// iota: 3, PST: -8
// 	PST
// )

// Skipping some values
const (
	// iota: 0, EST: -5
	EST Timezone = -(5 + iota)
	
	// _ is the blank identifier
    // iota: 1
	_

	// iota: 2, MST: -7
	MST 

	// iota: 3, PST: -8
	PST
)



// https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
func main() {
	fmt.Println("Enum Example from Medium")
	fmt.Println(Sunday)
	fmt.Println(Friday)

	fmt.Printf("Which day is it? %s\n", Sunday)

	fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.Weekend())
	fmt.Printf("Is %s a weekend day? %t\n", Monday.String(), Monday.Weekend())

}