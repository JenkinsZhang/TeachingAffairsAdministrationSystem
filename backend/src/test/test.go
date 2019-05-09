package test

import "fmt"

type user struct {
	id   string
	name string
}

func (u user) sing(song string) {
	fmt.Println(song)
}
type s interface{
	sing(song string)
}
func main() {
	me := user{"001","ybmj"}
	var i s
	i := me
	i.sing("love you")
}
