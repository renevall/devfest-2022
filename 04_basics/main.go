package main

import "fmt"

// Users representa la data que manejamos de un usuario.
type User struct {
	Name     string
	LastName string
	Phone    int
	Address  string
}

// Header retorna la lista de campos de la estructura User.
func (u User) Header() []string {
	return []string{"name", "last_name", "phone", "address"}
}

// SetName actualiza el nombre del usuario
func (u *User) SetName(name string) {
	u.Name = name
}

// Me piden que implemente un proceso que permita encodear una lista de usuarios
// en un CSV o XML

type Encoder interface {
	Encode([]User) ([]byte, error)
}

// CVSEncoder transforma una lista de usuarios en formato CSV
type CSVEncoder struct {
	OutoutFile string
}

func (e *CSVEncoder) Encode(data []User) ([]byte, error) {
	// imagina un proceso de encodeo
	output := []byte("ejemplo de csv")

	return output, nil
}

// XML encoder transforma un lista de usuarios en formato XML
type XMLEncoder struct {
	OutoutFile string
}

func (e *XMLEncoder) Encode(data []User) ([]byte, error) {
	// imagina un proceso de encodeo
	output := []byte("ejemplo de xml")

	return output, nil
}

func Main() {

	format := "csv"

	sample := []User{
		{Name: "Norberto", LastName: "Fernandez", Phone: 33456787, Address: "Granada"},
		{Name: "Ivan", LastName: "Ramirez", Phone: 33456787, Address: "Jipilto"},
		{Name: "Carlos", LastName: "Echeverria", Phone: 33456787, Address: "Rio Blanco"},
		{Name: "Amanda", LastName: "Escamilla", Phone: 33456787, Address: "Rivas"},
	}

	var encoder Encoder

	switch format {
	case "csv":
		encoder = &CSVEncoder{}
	case "xml":
		encoder = &XMLEncoder{}
	}

	output, err := encoder.Encode(sample)
	if err != nil {
		fmt.Println(output)
	}

}
