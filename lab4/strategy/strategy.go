package strategy

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретные стратегии
type CreditCardPayment struct {
	CardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{CardNumber: cardNumber}
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. кредитной картой %s", amount, c.CardNumber[len(c.CardNumber)-4:])
}

type PayPalPayment struct {
	Email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{Email: email}
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. через PayPal (%s)", amount, p.Email)
}

type CryptoPayment struct {
	WalletAddress string
}

func NewCryptoPayment(walletAddress string) *CryptoPayment {
	return &CryptoPayment{WalletAddress: walletAddress}
}

func (c *CryptoPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. криптовалютой (%s)", amount, c.WalletAddress[:8])
}

// Контекст
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) ProcessPayment(amount float64) string {
	if p.strategy == nil {
		return "Не выбран способ оплаты"
	}
	return p.strategy.Pay(amount)
}