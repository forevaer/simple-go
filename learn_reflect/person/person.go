package person

type Person struct {
	Name string
	Age int
	Gender string
}

func NewPerson(name string, gender string, age int) Person{
	return Person{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}
