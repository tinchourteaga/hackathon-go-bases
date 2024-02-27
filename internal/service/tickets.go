package service

import (
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) error
	// Read read a Ticket by id
	Read(id int) (Ticket, int, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) error
	// Delete delete a Ticket by id
	Delete(id int) error
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) error {
	_, _, err := b.Read(t.Id)

	if err == nil {
		return fmt.Errorf("error en la creacion del ticket con id %d: id ya existente", t.Id)
	}

	b.Tickets = append(b.Tickets, t)
	fmt.Printf("Se ha creado el ticket: %v\n", t)
	return nil
}

func (b *bookings) Read(id int) (Ticket, int, error) {
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, i, nil
		}
	}
	return Ticket{}, 0, fmt.Errorf("error en la lectura del ticket con id %d: id no existe", id)
}

func (b *bookings) Update(id int, t Ticket) error {
	_, index, err := b.Read(id)

	if err != nil {
		return err
	}

	b.Tickets[index] = t
	fmt.Printf("El ticket con id %d fue actualizado\n", id)

	return nil
}

func (b *bookings) Delete(id int) error {
	_, index, err := b.Read(id)

	if err != nil {
		return err
	}

	b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
	fmt.Printf("El ticket con id %d fue eliminado\n", id)

	return nil
}
