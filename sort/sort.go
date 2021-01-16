package main

import (
	"fmt"
  "sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type Interface interface{
  Len() int
  Less(i, j int) bool
  Swap(i, j int)
}

type byAge []user;
type byName []user;
type byLast []user;

func (a byAge) Len() int { return len(a); }
func (a byAge) Less(i, j int) bool {return a[i].Age < a[j].Age; }
func (a byAge) Swap(i, j int) { a[i], a[j] = a[j], a[i]; }

func (a byName) Len() int { return len(a); }
func (a byName) Less(i, j int) bool {return a[i].First < a[j].First; }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i]; }

func (a byLast) Len() int { return len(a); }
func (a byLast) Less(i, j int) bool {return a[i].Last < a[j].Last; }
func (a byLast) Swap(i, j int) { a[i], a[j] = a[j], a[i]; }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

  sort.Sort(byAge(users));
  for i, v := range users {
      fmt.Println("person ", i+1);
      fmt.Println("\tName : ", v.First);
      fmt.Println("\tLast : ", v.Last);
      fmt.Println("\tAge : ", v.Age);
      fmt.Println("\tSaying : ", v.Sayings);
  }

  sort.Sort(byName(users));
  for i, v := range users {
      fmt.Println("person ", i+1);
      fmt.Println("\tName : ", v.First);
      fmt.Println("\tLast : ", v.Last);
      fmt.Println("\tAge : ", v.Age);
      fmt.Println("\tSaying : ", v.Sayings);
  }

	sort.Sort(byLast(users));
  for i, v := range users {
      fmt.Println("person ", i+1);
      fmt.Println("\tName : ", v.First);
      fmt.Println("\tLast : ", v.Last);
      fmt.Println("\tAge : ", v.Age);
      fmt.Println("\tSaying : ", v.Sayings);
  }

}
