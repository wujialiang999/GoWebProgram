package cls1

import (
	"fmt"
	"goStudy/ch08/base"
)

type Class1 struct {
}

func (c *Class1) Do() {
	fmt.Println("class1")
}

func init() {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
