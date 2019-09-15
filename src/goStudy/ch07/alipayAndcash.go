package main

import "fmt"

type Alipay struct {
}

type Cash struct {
}

type CantainCanUseFaceID interface {
	CanUserFaceID()
}

type ContainStolen interface {
	Stolen()
}

func (a *Alipay) CanUserFaceID() {

}

func (c *Cash) Stolen() {

}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T can be stolen\n", payMethod)
	}
}

func main() {
	print(new(Alipay))
	print(new(Cash))
}
