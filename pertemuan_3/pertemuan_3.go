package pertemuan_3

import "fmt"

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (person *Person) SayHello() string {
	return "Hello my name is " + person.Name
}

func ChangeNamePerson(person *Person, newName string) {
	person.Name = newName
}

func init() {
	mahasiswa := map[string]string{
		"firstName": "Naufal",
		"lastName":  "Hakim",
		"hobby":     "Coding",
	}

	fmt.Println("=== Pertemuan 2 ===")
	fmt.Println(mahasiswa)

	person := NewPerson("Naufal")
	fmt.Println(person.SayHello())

	ChangeNamePerson(person, "Hakim")
	fmt.Println(person.SayHello())
}
