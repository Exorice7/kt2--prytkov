package animal

type Animal interface {
	Eat() string
	Sound() string
	Move() string
	Age() int
}

type Zebra struct {
	AgeValue int
}

func (z *Zebra) Sound() string {
	return "Зебра фыркает и ревет"
}

func (z *Zebra) Move() string {
	return "Зебра шагает"
}

func (z *Zebra) Age() int {
	return z.AgeValue
}

func (z *Zebra) Eat() string {
	return "Зебра ест траву"
}

type Tiger struct {
	AgeValue int
}

func (t *Tiger) Sound() string {
	return "Тигр рычит"
}

func (t *Tiger) Move() string {
	return "Тигр шагает"
}

func (t *Tiger) Age() int {
	return t.AgeValue
}

func (t *Tiger) Eat() string {
	return "Тигр ест животных"
}

type Panda struct {
	AgeValue int
}

func (p *Panda) Sound() string {
	return "Панда мычит"
}

func (p *Panda) Move() string {
	return "Панда шагает"
}

func (p *Panda) Age() int {
	return p.AgeValue
}

func (p *Panda) Eat() string {
	return "Панда ест бамбук"
}
