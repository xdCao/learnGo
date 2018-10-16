package main

type Player struct {
	curPos    Vec2
	targetPos Vec2
	speed     float32
}

func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

func (p *Player) Pos() (v Vec2) {
	return p.curPos
}

func (p *Player) IsArrived() bool {
	return p.curPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.curPos).Normalize()
		newPos := p.curPos.Add(dir.ScaleFloat(p.speed))
		p.curPos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
