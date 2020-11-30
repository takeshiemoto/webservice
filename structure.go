package structure

import "fmt"

func Structure()  {
	s := Shop{}
	s.setName("My Shop")
	fmt.Println(s.name)

	c := Cart{}
	c.Add("Book")
	c.Add("Desk")
	fmt.Println(c.items)
}
