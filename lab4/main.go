package main

import (
	"fmt"

	"github.com/lypolix/test-1/adapter"
	"github.com/lypolix/test-1/bridge"
	"github.com/lypolix/test-1/chain"
	"github.com/lypolix/test-1/iterator"
	"github.com/lypolix/test-1/proxy"
	"github.com/lypolix/test-1/strategy"
)

func main() {
	fmt.Println("=== Демонстрация поведенческих и структурных паттернов ===")

	// 1. Strategy Pattern
	fmt.Println("\n1. Strategy Pattern:")
	processor := &strategy.PaymentProcessor{}

	// Кредитная карта
	processor.SetStrategy(strategy.NewCreditCardPayment("1234567812345678"))
	fmt.Println(processor.ProcessPayment(1500.50))

	// PayPal
	processor.SetStrategy(strategy.NewPayPalPayment("user@example.com"))
	fmt.Println(processor.ProcessPayment(200.75))

	// Криптовалюта
	processor.SetStrategy(strategy.NewCryptoPayment("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"))
	fmt.Println(processor.ProcessPayment(300.25))

	// 2. Chain of Responsibility Pattern
	fmt.Println("\n2. Chain of Responsibility Pattern:")
	authHandler := &chain.AuthenticationHandler{}
	authzHandler := &chain.AuthorizationHandler{}
	validationHandler := &chain.ValidationHandler{}
	processingHandler := &chain.ProcessingHandler{}

	// Строим цепочку
	authHandler.SetNext(authzHandler).SetNext(validationHandler).SetNext(processingHandler)

	// Тестируем цепочку
	fmt.Println("Тест 1 - Успешный запрос:")
	result := authHandler.Handle("success")
	fmt.Printf("Результат: %s\n", result)

	fmt.Println("Тест 2 - Ошибка аутентификации:")
	result = authHandler.Handle("auth_error")
	fmt.Printf("Результат: %s\n", result)

	fmt.Println("Тест 3 - Ошибка авторизации:")
	result = authHandler.Handle("authz_error")
	fmt.Printf("Результат: %s\n", result)

	// 3. Iterator Pattern
	fmt.Println("\n3. Iterator Pattern:")
	library := &iterator.BookCollection{}
	library.AddBook("Война и мир")
	library.AddBook("Преступление и наказание")
	library.AddBook("Мастер и Маргарита")

	bookIterator := library.CreateIterator()
	fmt.Println("Книги в коллекции:")
	for bookIterator.HasNext() {
		book := bookIterator.Next()
		fmt.Printf(" - %s\n", book)
	}

	// 4. Proxy Pattern
	fmt.Println("\n4. Proxy Pattern:")
	image1 := proxy.NewImageProxy("photo1.jpg")
	image2 := proxy.NewImageProxy("photo2.jpg")

	// Изображение загрузится только при вызове Display
	fmt.Println(image1.Display())
	fmt.Println(image2.Display())

	// 5. Bridge Pattern
	fmt.Println("\n5. Bridge Pattern:")
	tv := bridge.NewTV()
	remote := bridge.NewRemoteControl(tv)

	remote.TogglePower() // Включить TV
	remote.VolumeUp()
	remote.VolumeUp()
	remote.VolumeDown()
	remote.TogglePower() // Выключить TV

	radio := bridge.NewRadio()
	advancedRemote := bridge.NewAdvancedRemoteControl(radio)

	advancedRemote.TogglePower() // Включить радио
	advancedRemote.VolumeUp()
	advancedRemote.Mute()
	advancedRemote.TogglePower() // Выключить радио

	// 6. Adapter Pattern
	fmt.Println("\n6. Adapter Pattern:")
	// Адаптер для принтера
	printerAdapter := adapter.NewPrinterAdapter("Hello World")
	fmt.Println(printerAdapter.PrintStored())

	// Адаптер для розетки
	socketAdapter := adapter.NewSocketAdapter()
	fmt.Println(socketAdapter.PlugIn())
}