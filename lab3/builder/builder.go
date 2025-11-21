package builder

type House struct {
    walls   string
    windows string
    doors   string
    roof    string
    garden  bool
    garage  bool
}

func (h *House) Show() {
    println("House details:")
    println("  Walls:", h.walls)
    println("  Windows:", h.windows)
    println("  Doors:", h.doors)
    println("  Roof:", h.roof)
    println("  Garden:", h.garden)
    println("  Garage:", h.garage)
}

// Интерфейс строителя
type HouseBuilder interface {
    BuildWalls() HouseBuilder
    BuildWindows() HouseBuilder
    BuildDoors() HouseBuilder
    BuildRoof() HouseBuilder
    BuildGarden() HouseBuilder
    BuildGarage() HouseBuilder
    GetHouse() *House
}

// Конкретный строитель
type ConcreteHouseBuilder struct {
    house *House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
    return &ConcreteHouseBuilder{house: &House{}}
}

func (b *ConcreteHouseBuilder) BuildWalls() HouseBuilder {
    b.house.walls = "Brick walls"
    return b
}

func (b *ConcreteHouseBuilder) BuildWindows() HouseBuilder {
    b.house.windows = "Glass windows"
    return b
}

func (b *ConcreteHouseBuilder) BuildDoors() HouseBuilder {
    b.house.doors = "Wooden doors"
    return b
}

func (b *ConcreteHouseBuilder) BuildRoof() HouseBuilder {
    b.house.roof = "Tile roof"
    return b
}

func (b *ConcreteHouseBuilder) BuildGarden() HouseBuilder {
    b.house.garden = true
    return b
}

func (b *ConcreteHouseBuilder) BuildGarage() HouseBuilder {
    b.house.garage = true
    return b
}

func (b *ConcreteHouseBuilder) GetHouse() *House {
    return b.house
}

// Директор
type Director struct {
    builder HouseBuilder
}

// NewDirector создает нового директора
func NewDirector(builder HouseBuilder) *Director {
    return &Director{builder: builder}
}

func (d *Director) BuildSimpleHouse() *House {
    return d.builder.
        BuildWalls().
        BuildWindows().
        BuildDoors().
        BuildRoof().
        GetHouse()
}

func (d *Director) BuildLuxuryHouse() *House {
    return d.builder.
        BuildWalls().
        BuildWindows().
        BuildDoors().
        BuildRoof().
        BuildGarden().
        BuildGarage().
        GetHouse()
}