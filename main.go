package main

import (
	"case-transport/transport"

	"case-transport/route"
)

func main() {
	bus := transport.Bus{}
	train := transport.Train{}
	plane := transport.Plane{}

	travelRoute := route.Route{}
	travelRoute.AddVehicle(bus)
	travelRoute.AddVehicle(train)
	travelRoute.AddVehicle(plane)

	travelRoute.ShowVehicles()
	travelRoute.Travel()
}
