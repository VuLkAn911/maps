package main

import (
	"fmt"
)

func mapsEqual(a, b map[string]string) bool {
	if a, b := len(a), len(b); a != b {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}

	}
	return true
}

func compare(currentConfig, defaultConfig map[string]string) {
	if mapsEqual(currentConfig, defaultConfig) {
		fmt.Println("Конфигурации идентичны")
	} else {
		fmt.Println("Конфигурации не идентичны")
	}
}
func main() {
	// sub 1
	toolUsage := map[string]int{"Go": 3, "VSCode": 5, "Git": 2}

	for i, v := range toolUsage {
		fmt.Println(i, v)
	}

	// sub 2

	buildStatus := map[string]bool{"Build": true, "run": false}

	if buildStatus["Build"] == true {
		fmt.Println("Сборка прошла успешно")

	} else {
		fmt.Println("Сборка не прошла успешно")
	}

	// sub 3

	// var name string
	// fmt.Printf("Введите ваше имя:")
	// fmt.Scan(&name)

	// userInfo := make(map[string]string)
	// userInfo["name"] = name
	// userInfo["isLoggedIn"] = "true"

	// fmt.Printf("Пользователь %s вошёл в систему", userInfo["name"])

	// sub 4
	max := 0
	cpuLoad := map[int]int{1: 40, 2: 65, 3: 30}

	for _, v := range cpuLoad {

		if v > max {
			max = v
		}

	}
	fmt.Println(max)

	// sub 5

	examResults := map[string]int{"Aruzhan": 85, "Dias": 92, "Alina": 78}

	for i, v := range examResults {

		if v >= 80 {
			fmt.Println("Ученик", i, "cдал(a) экзамен")
		} else {
			fmt.Println(i, "Не сдал(a) экзамен")
		}
	}

	// sub 6

	words := []string{"go", "is", "fast"}

	worldLength := make(map[string]int)

	for _, v := range words {

		worldLength[v] = len(v)

	}
	fmt.Println(worldLength)

	// sub 7

	// var nameMenu string

	// fmt.Println("Введите название блюда:")
	// fmt.Scan(&nameMenu)

	// menu := map[string]int{"Burger": 2500, "Pizza": 3200, "Tea": 500}

	// if _, ok := menu[nameMenu]; ok {
	// 	fmt.Println(menu[nameMenu])
	// } else {
	// 	fmt.Println("Такого блюда нет в меню")
	// }

	// sub 8

	loginAttempts := make(map[string]int)

	loginAttempts["admin"] = 2
	loginAttempts["guest"] = 0

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ Заблокирован")
	}

	// sub 9
	total := make(map[int]int)

	sales := [2][3]int{{10, 20, 30}, {40, 50, 60}}

	for i, row := range sales {
		summ := 0
		for _, v := range row {
			summ += v

		}
		total[i+1] = summ

	}

	fmt.Println(total)

	// sub 10

	numbers := []int{4, 7, 2, 9, 5}

	numberStatus := make(map[int]string)

	for _, v := range numbers {
		if v%2 == 0 {
			numberStatus[v] = "even"
		} else {
			numberStatus[v] = "odd"

		}
	}

	fmt.Println(numberStatus)

	// sub 11

	defaultConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}

	currentConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}

	compare(defaultConfig, currentConfig)

	currentConfig["mode"] = "debug"

	compare(defaultConfig, currentConfig)
}
