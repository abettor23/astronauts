package main

import (
	"fmt"
)

func main() {
	fmt.Println("Для миссии на Марс отобрали трёх лучших космонавтов.")
	fmt.Println("Но выбрать командира сложно, выберем по наивысшему IQ.")

	fmt.Println("Введите IQ первого космонавта:")
	var astronautOneIQ int
	fmt.Scan(&astronautOneIQ)

	fmt.Println("Введите IQ второго космонавта:")
	var astronautTwoIQ int
	fmt.Scan(&astronautTwoIQ)

	fmt.Println("Введите IQ третьего космонавта:")
	var astronautThreeIQ int
	fmt.Scan(&astronautThreeIQ)

	var highestIQ int

	if astronautOneIQ > astronautTwoIQ {
		if astronautOneIQ > astronautThreeIQ {
			highestIQ = astronautOneIQ
			fmt.Print("Наивысший IQ - ", highestIQ, ". Командиром выбран первый космоноват!")
		}
	} else if astronautTwoIQ > astronautThreeIQ {
		highestIQ = astronautTwoIQ
		fmt.Print("Наивысший IQ - ", highestIQ, ". Командиром выбран второй космоноват!")
	} else if astronautThreeIQ > astronautTwoIQ {
		highestIQ = astronautThreeIQ
		fmt.Print("Наивысший IQ - ", highestIQ, ". Командиром выбран третий космоноват!")
	}
}
