package abstractfactory

// Интерфейсы продуктов
type Button interface {
    Render()
    OnClick()
}

type Checkbox interface {
    Render()
    IsChecked() bool
}

// Конкретные продукты для Windows
type WindowsButton struct{}

func (w *WindowsButton) Render() {
    println("Rendering Windows style button")
}

func (w *WindowsButton) OnClick() {
    println("Windows button clicked")
}

type WindowsCheckbox struct {
    checked bool
}

func (w *WindowsCheckbox) Render() {
    println("Rendering Windows style checkbox")
}

func (w *WindowsCheckbox) IsChecked() bool {
    return w.checked
}

// Конкретные продукты для Mac
type MacButton struct{}

func (m *MacButton) Render() {
    println("Rendering Mac style button")
}

func (m *MacButton) OnClick() {
    println("Mac button clicked")
}

type MacCheckbox struct {
    checked bool
}

func (m *MacCheckbox) Render() {
    println("Rendering Mac style checkbox")
}

func (m *MacCheckbox) IsChecked() bool {
    return m.checked
}

// Абстрактная фабрика
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// Конкретные фабрики
type WindowsFactory struct{}

func (w *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

func (w *WindowsFactory) CreateCheckbox() Checkbox {
    return &WindowsCheckbox{checked: false}
}

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (m *MacFactory) CreateCheckbox() Checkbox {
    return &MacCheckbox{checked: true}
}

// Фабрика фабрик
func GetFactory(os string) GUIFactory {
    switch os {
    case "windows":
        return &WindowsFactory{}
    case "mac":
        return &MacFactory{}
    default:
        return nil
    }
}