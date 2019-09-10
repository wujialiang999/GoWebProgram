package main

import (
	"math"
)

type Vec struct {
	X, Y float32
}

func (v Vec) Add(other Vec) Vec {
	return Vec{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v Vec) Scale(s float32) Vec {
	return Vec{
		v.X * s,
		v.Y * s,
	}
}

func (v Vec) DistanceTo(other Vec) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dis := math.Sqrt(float64(dx*dx + dy*dy))
	result := float32(dis)
	return result
}

func (v Vec) Normailze() Vec {
	msg := v.X*v.X + v.Y*v.Y
	if msg > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(msg)))

		return Vec{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec{0, 0}
}

//二维矢量模拟玩家移动
