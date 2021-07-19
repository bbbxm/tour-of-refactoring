package chapter1

type invoice struct {
	Customer     string
	Performances []performance
}

type performance struct {
	PlayID   string
	Audience int
	Play     play
}

type play struct {
	Name string
	Type string
}
