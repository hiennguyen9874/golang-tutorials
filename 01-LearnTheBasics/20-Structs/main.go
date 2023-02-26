package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
	info StudentInfo
}

type Student2 struct {
	id   int
	name string
	info map[int]int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	var st3 = struct {
		id   int
		name string
	}{123, "Hien"}
	fmt.Println(st3)

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	st4 := Student{1, "Hien", StudentInfo{"A", "a@gmail.com", 27}}
	fmt.Println(st4)

	p1 := person{"A", 1}
	p2 := person{"A", 1}
	fmt.Println(p1 == p2)

	st5 := Student2{1, "A", map[int]int{1: 1}}
	st6 := Student2{1, "A", map[int]int{1: 1}}
	fmt.Println(st5 == st6)

}
