package bridge

import "fmt"

// Интерфейс реализации
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(volume int)
}

// Конкретные реализации
type TV struct {
	on     bool
	volume int
}

func NewTV() *TV {
	return &TV{}
}

func (t *TV) IsEnabled() bool {
	return t.on
}

func (t *TV) Enable() {
	t.on = true
	fmt.Println("TV включен")
}

func (t *TV) Disable() {
	t.on = false
	fmt.Println("TV выключен")
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
	fmt.Printf("Громкость TV установлена на %d\n", volume)
}

type Radio struct {
	on     bool
	volume int
}

func NewRadio() *Radio {
	return &Radio{}
}

func (r *Radio) IsEnabled() bool {
	return r.on
}

func (r *Radio) Enable() {
	r.on = true
	fmt.Println("Радио включено")
}

func (r *Radio) Disable() {
	r.on = false
	fmt.Println("Радио выключено")
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Printf("Громкость радио установлена на %d\n", volume)
}

// Абстракция
type RemoteControl struct {
	device Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

func (r *RemoteControl) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *RemoteControl) VolumeDown() {
	r.device.SetVolume(r.device.GetVolume() - 10)
}

func (r *RemoteControl) VolumeUp() {
	r.device.SetVolume(r.device.GetVolume() + 10)
}

// Расширенная абстракция
type AdvancedRemoteControl struct {
	RemoteControl
}

func NewAdvancedRemoteControl(device Device) *AdvancedRemoteControl {
	return &AdvancedRemoteControl{
		RemoteControl: RemoteControl{device: device},
	}
}

func (a *AdvancedRemoteControl) Mute() {
	a.device.SetVolume(0)
	fmt.Println("Устройство заглушено")
}
