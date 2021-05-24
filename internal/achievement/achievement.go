package achievement

import "fmt"

type Achievement struct {
	Id          uint
	Name        string
	Description string
}

func New() *Achievement {
	return &Achievement{
		Id:          0,
		Name:        "",
		Description: "",
	}
}

func MustNew() Achievement {
	return Achievement{
		Id:          0,
		Name:        "",
		Description: "",
	}
}

func (a *Achievement) Init(id uint, name string, description string) {
	a.Id = id
	a.Name = name
	a.Description = description
}

func (a *Achievement) Close() {
	fmt.Printf("Achievement '%q' closed\n", a)
}

func (a *Achievement) String() string {
	return fmt.Sprintf("Achievement (id=%d, name is %q, description is %q", a.Id, a.Name, a.Description)
}
