package particle

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	RECT_WIDTH  = 10
	RECT_HEIGHT = 10
)

type Particle struct {
	Energy     int
	Xpos, Ypos int
	Rect       sdl.Rect
	Color      [4]int
}

func NewParticle(energy int, color [4]int) *Particle {
	return &Particle{
		Energy: energy,
		Color:  color,
	}
}

func (p *Particle) SetPosition(x, y int) {
	p.Xpos = x
	p.Ypos = y

	p.Rect.X = int32(x) - RECT_WIDTH/2
	p.Rect.Y = int32(y) - RECT_HEIGHT/2
}

func (p *Particle) SetRandomPosition(winWidth, winHeight int) {
	rand.Seed(time.Now().UnixNano())
	p.SetPosition(rand.Intn(winWidth), rand.Intn(winHeight))
}