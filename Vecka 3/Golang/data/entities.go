package data

type Employee struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Name string `json:"name"`
}

func (emp Employee) IsTrash() bool {
	if emp.Name == "David" {
		return true
	}

	return false
}