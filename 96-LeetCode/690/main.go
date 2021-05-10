package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	emap := make(map[int]*Employee)
	for _, _i := range employees {
		emap[_i.Id] = _i
	}
	sum := 0
	var rang func(int)
	rang = func(idx int) {
		e := emap[idx]
		sum += e.Importance
		if len(e.Subordinates) > 0 {
			for _, _id := range e.Subordinates {
				rang(_id)
			}
		}
	}
	rang(id)
	return sum
}
