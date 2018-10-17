package base

type Class interface {
	Do()
}

var (
	facBYName = make(map[string]func() Class)
)

func Register(name string, fac func() Class) {
	facBYName[name] = fac
}

func Create(name string) Class {
	if f, ok := facBYName[name]; ok {
		return f()
	} else {
		panic("name not found ")
	}
}
