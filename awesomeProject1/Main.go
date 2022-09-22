package main

func main() {
	person := Person{"Bob"}
	hhKz := JobSite{}

	hhKz.addVacancy("Senior HTML Developer")
	hhKz.subscribe(person)
	hhKz.addVacancy("Java Junior Developer")

	hhKz.removeVacancy("Senior HTML Developer")
}
