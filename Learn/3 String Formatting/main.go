package main

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
// => Happy birthday Frank! You are now 58 years old!
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.

// AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// =>
// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s", name, table, direction, distance, neighbor)
	panic("Please implement the AssignTable function")
}

func main() {
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
