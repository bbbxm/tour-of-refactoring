package chapter1

type invoice struct {
	Customer     string
	Performances []performance
}

type performance struct {
	PlayID   string
	Audience int
}

type play struct {
	Name string
	Type string
}
