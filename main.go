package main

import (
	"fmt"
	"github.com/ngzijian/golib/collection"
	"github.com/ngzijian/golib/math"
)

func main() {
	arr := []string{"1", "2", "1", "2", "3"}
	list := collection.List{
		Val: arr,
	}
	list.Distinct()
	fmt.Println(list.Val)

	ma := math.Math{}
	r := ma.Add(2, 3)

	fmt.Println(r)
}
