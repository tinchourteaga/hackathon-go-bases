package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file := file.File{Path: "tickets.csv"}
	tickets, err := file.Read()

	if err != nil {
		// Quiero un panic aca porque si no puedo traerme el archivo, no puedo trabajar
		panic(err)
	}
	// Funcion para obtener tickets del archivo csv
	bookings := service.NewBookings(tickets)
	err = bookings.Create(service.Ticket{Id: 1001, Names: "pepe", Email: "pepe@gmail.com", Destination: "Argentina", Date: "11:23", Price: 500})
	err = bookings.Create(service.Ticket{Id: 1002, Names: "pepa", Email: "pepa@gmail.com", Destination: "Brasil", Date: "11:24", Price: 800})
	err = bookings.Create(service.Ticket{Id: 2, Names: "pepo", Email: "pepo@gmail.com", Destination: "Uruguay", Date: "11:25", Price: 600})

	if err != nil {
		fmt.Println(err)
	}

	bookings.Update(3, service.Ticket{Id: 3, Names: "Martin", Email: "martin@gmail.com", Destination: "Cancun", Date: "10:50", Price: 5000})
	bookings.Delete(3)
}
