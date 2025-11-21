package adapter

import "fmt"

// Старая система (адаптируемый класс)
type LegacyPrinter struct{}

func (l *LegacyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s", s)
}

// Новый интерфейс
type ModernPrinter interface {
	PrintStored() string
}

// Адаптер
type PrinterAdapter struct {
	OldPrinter *LegacyPrinter
	Msg        string
}

func NewPrinterAdapter(msg string) *PrinterAdapter {
	return &PrinterAdapter{
		OldPrinter: &LegacyPrinter{},
		Msg:        msg,
	}
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(p.Msg)
	}
	return p.Msg
}

// Другой пример адаптера
type EuropeanSocket interface {
	PlugIn() string
}

type AmericanSocket struct{}

func (a *AmericanSocket) PlugInAmerican() string {
	return "American plug connected"
}

type SocketAdapter struct {
	AmericanSocket *AmericanSocket
}

func NewSocketAdapter() *SocketAdapter {
	return &SocketAdapter{
		AmericanSocket: &AmericanSocket{},
	}
}

func (s *SocketAdapter) PlugIn() string {
	return s.AmericanSocket.PlugInAmerican() + " through adapter"
}