package main

import "fmt"

type Flyer interface {
	Fly()
}
type Walker interface {
	Walk()
}
type Bird struct {
}

func (b *Bird) Walk() {
	fmt.Println("bird:walk")
}

func (b *Bird) Fly() {
	fmt.Println("bird:fly")
}

type Pig struct {
}

func (p *Pig) Walk() {
	fmt.Println("pig:walk")
}
func main() {
	//将接口转换为其他接口
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}
	for name, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)
		fmt.Printf("name:%s isFlyer:%v isWalker:%v\n", name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}
	//将接口转换为其他类型
	p1 := new(Pig)
	var a Walker = p1 //pig隐式转换为Walker接口
	p2 := a.(*Pig)    //walker转换为pig接口
	fmt.Printf("p1:%p,p2:%p\n", p1, p2)
}

/*
name:bird isFlyer:true isWalker:true
bird:fly
bird:walk
name:pig isFlyer:false isWalker:true
pig:walk
*/
