package factory

type Transport interface {
    Deliver(destination string)
    GetName() string
}

type Truck struct {
    name string
}

func (t *Truck) Deliver(destination string) {
    println("Грузовик", t.name, "доставляет груз в", destination, "по земле")
}

func (t *Truck) GetName() string {
    return t.name
}

type Ship struct {
    name string
}

func (s *Ship) Deliver(destination string) {
    println("Корабль", s.name, "доставляет груз в", destination, "по морю")
}

func (s *Ship) GetName() string {
    return s.name
}

type Logistics interface {
    PlanDelivery(destination string) Transport
}

type RoadLogistics struct{}

func (rl *RoadLogistics) PlanDelivery(destination string) Transport {
    return &Truck{name: "Volvo Truck"}
}

type SeaLogistics struct{}

func (sl *SeaLogistics) PlanDelivery(destination string) Transport {
    return &Ship{name: "Cargo Ship"}
}