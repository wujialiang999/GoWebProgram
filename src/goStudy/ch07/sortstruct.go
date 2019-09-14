package main

import (
	"fmt"
	"sort"
)

type HeroKind int

//定义HeroKind常量，类似枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) Len() int {
	return len(s)
}

func (s Heros) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	heros2 := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
		&Hero{"貂蝉", Mage},
	}
	//需要写多个接口实现
	sort.Sort(heros2)
	//简单方法
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
		&Hero{"貂蝉", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros2 {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("")
	for _, v := range heros {
		fmt.Printf("%v\n", v)
	}
}
