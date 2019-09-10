package main

import ()

type Player struct {
	currPos   Vec
	targetPos Vec
	speed     float32
}

func (p *Player) MoveTo(v Vec) {
	p.targetPos = v
}
func (p *Player) Pos() Vec {
	return p.currPos
}

func (p *Player) IsArrived() bool {
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) Update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normailze()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
