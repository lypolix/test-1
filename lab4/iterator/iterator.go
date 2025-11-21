package iterator

// Интерфейс итератора
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Интерфейс коллекции
type Collection interface {
	CreateIterator() Iterator
}

// Конкретная коллекция
type BookCollection struct {
	books []string
}

func (b *BookCollection) AddBook(book string) {
	b.books = append(b.books, book)
}

func (b *BookCollection) CreateIterator() Iterator {
	return &BookIterator{
		books: b.books,
		index: 0,
	}
}

// Конкретный итератор
type BookIterator struct {
	books []string
	index int
}

func (i *BookIterator) HasNext() bool {
	return i.index < len(i.books)
}

func (i *BookIterator) Next() interface{} {
	if i.HasNext() {
		book := i.books[i.index]
		i.index++
		return book
	}
	return nil
}