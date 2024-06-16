package route

import (
	"fmt"
	"case-transport/transport"
)

// Route структура для маршруту
type Route struct {
	vehicles []transport.PublicTransport
}

// AddVehicle додає транспортний засіб до маршруту
func (r *Route) AddVehicle(vehicle transport.PublicTransport) {
	r.vehicles = append(r.vehicles, vehicle)
}

// ShowVehicles показывує список транспортних засобів на маршруті
func (r *Route) ShowVehicles() {
	for i, vehicle := range r.vehicles {
		fmt.Printf("Транспортний засіб %d: %T\n", i+1, vehicle)
	}
}

// Travel проходить по маршруту и виводить процес на екран
func (r *Route) Travel() {
	for _, vehicle := range r.vehicles {
		vehicle.BoardPassengers()
		vehicle.UnboardPassengers()
	}
}
