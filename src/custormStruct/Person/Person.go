package Person

type Person struct {
	Title string
	Forenames []string
	surname string
}

type Author1 struct {
	Names Person
	TiTle []string
	YearBorn int
}

type Author2 struct {
	Person
	TiTle []string
	YearBorn int
}

//author1.Names.TiTle
//author2.Forenames
//author2.TiTle = []string{}
//anthor2.Person.TiTle

type Count struct {
	Id int
}

func (c *Count) Incre() {
	c.Id++
}

type Tasks struct {
	slice []string
	Count
}

func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
	tasks.Incre() // = tasks.Count.Incre()
}