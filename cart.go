package structure

type Cart struct {
	items []string
}

func (c *Cart) Add(item string) {
	c.items = append(c.items, item)
}