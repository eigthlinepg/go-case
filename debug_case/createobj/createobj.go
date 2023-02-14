package main

type User struct {
	Name string
	Age  int
}

//go:noinline
func main() {

	var i int

	var in *int = new(int)

	su := new(User)
	su.Name = "test"
	su.Age = 10

	println(&i, in)
	println(su, su.Name, su.Age)

}
