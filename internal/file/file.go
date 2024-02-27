package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	tickets := []service.Ticket{}
	file, err := os.Open(f.Path)

	if err != nil {
		file.Close()
		return nil, errors.New("error en la apertura del archivo")
	}

	reader := csv.NewReader(file)
	for {
		content, err := reader.Read() // Trae de a una linea como un array de strings

		if err == io.EOF {
			break
		}

		if err != nil {
			file.Close()
			return nil, errors.New("error en la lectura del archivo")
		}

		id, _ := strconv.Atoi(content[0])
		price, _ := strconv.Atoi(content[5])
		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       content[1],
			Email:       content[2],
			Destination: content[3],
			Date:        content[4],
			Price:       price,
		})
	}
	file.Close()
	return tickets, nil
}

func (f *File) Write(ticket service.Ticket) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		file.Close()
		return errors.New("error en la apertura del archivo")
	}

	data := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	file.Write([]byte(data))

	if err != nil {
		file.Close()
		return errors.New("error en la escritura del archivo")
	}
	file.Close()
	return nil
}
