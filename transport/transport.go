package transport

import "fmt"

// Bus структура для автобуса
type Bus struct{}

func (b Bus) BoardPassengers() {
	fmt.Println("Автобус приймає пасажирів.")
}

func (b Bus) UnboardPassengers() {
	fmt.Println("Автобус висаджує пасажирів.")
}

// Train структура для поїзду
type Train struct{}

func (t Train) BoardPassengers() {
	fmt.Println("Потяг приймає пасажирів.")
}

func (t Train) UnboardPassengers() {
	fmt.Println("Потяг висаджує пасажирів.")
}

// Plane структура для літака
type Plane struct{}

func (p Plane) BoardPassengers() {
	fmt.Println("Літак приймає пасажирів.")
}

func (p Plane) UnboardPassengers() {
	fmt.Println("Літак высаджує пасажирів.")
}
