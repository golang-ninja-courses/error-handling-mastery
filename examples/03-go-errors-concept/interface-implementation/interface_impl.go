package interfaceimpl

type Ducker interface {
	Quack() string // крякать
	Swim()         // плавать
}

type Duck struct{} // Утка "реализует" интерфейс Ducker.

func (d Duck) Quack() string {
	return "Quack!"
}

func (d Duck) Swim() {
	// Притворимся, что здесь утка плавает.
}

type Otter struct{} // Выдра Otter не "реализует" интерфейс Ducker, потому что не умеет крякать.

func (o Otter) Squeak() string {
	return "Squaek!"
}
