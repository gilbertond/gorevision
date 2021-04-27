package main

type Job struct {
	name        string
	description string
	id          int
}

func main() {
	// j1 := new(Job)
	// j2 := new(Job)
	
	data := struct {
		string
		Job
	}{"gilbert ",
		*new(Job)}
	
	
}
