package main

import (
    "fmt"
    "sync"

    "github.com/lypolix/test-1/abstractfactory"
    "github.com/lypolix/test-1/builder"
    "github.com/lypolix/test-1/factory"
    "github.com/lypolix/test-1/singleton"
)

func main() {
    fmt.Println("=== Демонстрация порождающих паттернов ===\n")

    // 1. Singleton
    fmt.Println("1. Singleton Pattern:")
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            db := singleton.GetInstance()
            db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %d", id))
        }(i)
    }
    wg.Wait()
    fmt.Println()

    // 2. Factory Method
    fmt.Println("2. Factory Method Pattern:")
    roadLogistics := &factory.RoadLogistics{}
    seaLogistics := &factory.SeaLogistics{}
    
    truck := roadLogistics.PlanDelivery("Москва")
    ship := seaLogistics.PlanDelivery("Нью-Йорк")
    
    truck.Deliver("Москва")
    ship.Deliver("Нью-Йорк")
    fmt.Println()

    // 3. Abstract Factory
    fmt.Println("3. Abstract Factory Pattern:")
    windowsFactory := abstractfactory.GetFactory("windows")
    macFactory := abstractfactory.GetFactory("mac")
    
    winButton := windowsFactory.CreateButton()
    winCheckbox := windowsFactory.CreateCheckbox()
    
    macButton := macFactory.CreateButton()
    macCheckbox := macFactory.CreateCheckbox()
    
    winButton.Render()
    winButton.OnClick()
    winCheckbox.Render()
    fmt.Printf("Windows checkbox checked: %t\n\n", winCheckbox.IsChecked())
    
    macButton.Render()
    macButton.OnClick()
    macCheckbox.Render()
    fmt.Printf("Mac checkbox checked: %t\n\n", macCheckbox.IsChecked())

    // 4. Builder
    fmt.Println("4. Builder Pattern:")
    houseBuilder := builder.NewConcreteHouseBuilder()
    director := builder.NewDirector(houseBuilder)
    
    simpleHouse := director.BuildSimpleHouse()
    fmt.Println("Simple House:")
    simpleHouse.Show()
    fmt.Println()
    
    luxuryHouse := director.BuildLuxuryHouse()
    fmt.Println("Luxury House:")
    luxuryHouse.Show()
}