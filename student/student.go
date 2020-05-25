package student

type Student struct {
	Id int
	name string
}

func CreateStudent(id int, name string) Student {
	return Student{id,name}
}