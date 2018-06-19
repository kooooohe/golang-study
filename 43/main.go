package main

type User struct {
	Id   int
	Name string
}

func main() {
	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "TOkyo",
		{Id: 2, Name: "Jaro"}: "Osaka",
	}
	m2 := map[int]User{
		1: User{Id: 11, Name: "aa"},
	}
	m3 := map[int]map[int]string{
		1: {1: "Apple", 2: "Banana"},
	}
}
