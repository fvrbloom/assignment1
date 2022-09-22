package main

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (jobs *JobSite) subscribe(observer Observer) {
	//TODO implement me
	jobs.subscribers = append(jobs.subscribers, observer)
}

func (jobs *JobSite) unsubscribe(observer Observer) {
	//TODO implement me
	for m, n := range jobs.subscribers {
		if n == observer {
			jobs.subscribers = append(jobs.subscribers[:m], jobs.subscribers[m+1:]...)
		}
	}
}

func (jobs JobSite) sendAll() {
	//TODO implement me
	for _, subscriber := range jobs.subscribers {
		subscriber.handleEvent(jobs.vacancies)
	}
}
func (jobs *JobSite) addVacancy(vacancy string) {
	jobs.vacancies = append(jobs.vacancies, vacancy)
	jobs.sendAll()
}

func (jobs *JobSite) removeVacancy(vacancy string) {
	for m, n := range jobs.vacancies {
		if n == vacancy {
			jobs.vacancies = append(jobs.vacancies[:m], jobs.vacancies[m+1:]...)
		}
	}

	jobs.sendAll()
}
