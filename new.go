package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var finalBill int

func main() {
	fmt.Print("C: coffee, 50rs\n D: dosa, 90 rs\n T: tomato soup, 30rs\n I : idli , 30rs\n V : vada, 35rs.\n B: bhature &chhole, 40rs\n P: paneer pakoda, 40rs\n M: manchurian, 100rs\n H: hakka noodle, 80rs.\n F: French fries, 40rs\n J: jalebi, 40rs\n L: lemonade, 25rs\n S: spring roll, 30rs\n")
	OpenFunc()
}
func OpenFunc() {

	fmt.Print("Please enter order If Day Ends type END :")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	result := comFunc(line)
	st := line
	if result == 1 {
		fmt.Println("Day complete Final Bill :")
		Display()
	} else {
		BillCount(st)
		OpenFunc()
	}
}
func comFunc(s string) int {
	res := strings.Compare(s, "END")
	return res
}
func BillCount(s string) {
	var Bill int
	order := strings.ToUpper(string(s[2]))
	m := map[string]int{
		"C": 50,
		"D": 90,
		"T": 30,
		"I": 30,
		"V": 35,
		"B": 40,
		"P": 40,
		"M": 100,
		"H": 80,
		"F": 40,
		"J": 40,
		"L": 25,
		"S": 30,
	}
	v, _ := strconv.Atoi(string(s[0]))
	Bill = v * m[order]
	FinalBill(Bill)
}
func FinalBill(n int) {
	finalBill += n
}
func Display() {
	fmt.Print(finalBill)
}
