package chain

import "fmt"

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// Базовая структура обработчика
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) string {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return ""
}

// Конкретные обработчики
type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) Handle(request string) string {
	if request == "auth_error" {
		return "Ошибка аутентификации: неверные учетные данные"
	}
	fmt.Println("Аутентификация пройдена успешно")
	return a.BaseHandler.Handle(request)
}

type AuthorizationHandler struct {
	BaseHandler
}

func (a *AuthorizationHandler) Handle(request string) string {
	if request == "authz_error" {
		return "Ошибка авторизации: недостаточно прав"
	}
	fmt.Println("Авторизация пройдена успешно")
	return a.BaseHandler.Handle(request)
}

type ValidationHandler struct {
	BaseHandler
}

func (v *ValidationHandler) Handle(request string) string {
	if request == "validation_error" {
		return "Ошибка валидации: неверные данные"
	}
	fmt.Println("Валидация пройдена успешно")
	return v.BaseHandler.Handle(request)
}

type ProcessingHandler struct {
	BaseHandler
}

func (p *ProcessingHandler) Handle(request string) string {
	fmt.Println("Обработка запроса завершена успешно")
	return "Запрос обработан успешно"
}