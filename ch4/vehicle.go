package main

type Vehicle interface {
	RefillFuel()
}

type Car interface {
	Vehicle // embed interface
}

type Honda struct{}

func (h *Honda) RefillFuel() {

}

type Tesla struct{}

func (h *Tesla) RefillFuel() {
}

type YamahaMotorcycle struct {
}
