package main

func loadData() Class {
	return Class{
		Grade: 4,
		Teacher: Person{

			FirstName: "Herbert",
			LastName:  "Garrison",
			Age:       30,
		},
		Students: []Person{
			{
				FirstName: "Eric",
				LastName:  "Cartman",
				Age:       8,
			},
			{
				FirstName: "Stan",
				LastName:  "Marsh",
				Age:       8,
			},
			{
				FirstName: "Kyle",
				LastName:  "Broflovski",
				Age:       8,
			},
			{
				FirstName: "Kenny",
				LastName:  "McKormick",
				Age:       8,
			},
		},
	}
}
