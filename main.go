package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := task()
	w := bufio.NewWriter(os.Stdout)
	for _, elem := range ans {
		w.WriteString(elem)
		w.WriteString("\n")
	}
	w.Flush()
}

func task() []string {
	ans := make([]string, 0, 10) // Примерное capacity, чтобы не производить лишнее расширение
	recursion(8, "9", &ans)      // Вызов рекурсивной функции
	return ans
}

func recursion(i int, line string, ans *[]string) {
	// Худшее работы O(10), что является констатным временем
	num := getNum(line)

	if i == -1 {
		if num == 200 {
			*ans = append(*ans, line+"=200")
		}
		return
	}

	char := strconv.Itoa(i)
	recursion(i-1, line+"+"+char, ans)
	recursion(i-1, line+"-"+char, ans)
	recursion(i-1, line+char, ans)
}

func getNum(line string) int {
	terms := strings.Split(line, "+")
	sum := 0
	for _, term := range terms {
		if strings.Contains(term, "-") {
			subTerms := strings.Split(term, "-")
			subSum := 0
			for i, subTerm := range subTerms {
				val, _ := strconv.Atoi(subTerm)
				if i == 0 {
					subSum += val
				} else {
					subSum -= val
				}
			}
			sum += subSum
		} else {
			val, _ := strconv.Atoi(term)
			sum += val
		}
	}
	return sum
}
