package main

import "fmt"

func main(){
	hhkz :=JobSite{
		subscribers: nil,
		vacancies:   nil,
	}
	bob := Person{name:"BOB"}

	hhkz.subscribe(&bob)
	hhkz.addVaccacnies("Java Junior")

}


type Observer interface {
	handleEvent([]string)
}

type Person struct {
	name string
}

func  newPerson(name string)  *Person{
	p := new(Person)
	p.name = name
	return p
}
func (p * Person)handleEvent(vacancies []string){
	fmt.Println("HI dear " , p.name)
	fmt.Println("Vacancies update: ")
	for _,v := range vacancies{
		fmt.Println(v)
	}
}
type Observable interface {
	subscribe(Observer)
	unsubscribe(Observer)
	sendAll()
}

type JobSite struct {
	subscribers []Observer
	vacancies []string
}



func newJobSite(site JobSite)  *JobSite{
	j := new(JobSite)
	j.subscribers = make([]Observer, 0, 5)
	j.vacancies = make([]string, 0, 5)
	return j
}
func (j *JobSite) sendAll(){
	for _, s := range j.subscribers {
		s.handleEvent(j.vacancies)
	}
}

func (j * JobSite) subscribe(observer Observer){
	j.subscribers = append(j.subscribers, observer)
}

func (j * JobSite)unsubscribe(observer Observer){
	fmt.Println( observer,"unsubscribed")
}
func (j *JobSite) addVaccacnies(vaccancy string)  {
	j.vacancies = append(j.vacancies, vaccancy)
	j.sendAll()
}
func (j *JobSite) removeVaccancy(vaccancy string){
	fmt.Println(vaccancy, "removed")
}


