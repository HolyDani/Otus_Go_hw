package book

type Book struct {
	id     uint
	year   uint
	size   uint
	rate   float32
	title  string
	author string
}

func NewBook(newID, newYear, newSize uint, newRate float32, newAuthor, newTitle string) *Book {
	return &Book{
		id:     newID,
		year:   newYear,
		size:   newSize,
		rate:   newRate,
		author: newAuthor,
		title:  newTitle,
	}
}

func (b *Book) ID() *uint {
	return &b.id
}

func (b *Book) Year() *uint {
	return &b.year
}

func (b *Book) Size() *uint {
	return &b.size
}

func (b *Book) Rate() *float32 {
	return &b.rate
}

func (b *Book) Title() *string {
	return &b.title
}

func (b *Book) Author() *string {
	return &b.author
}

func (b *Book) SetID(id uint) {
	b.id = id
}

func (b *Book) SetSize(size uint) {
	b.size = size
}

func (b *Book) SetYear(year uint) {
	b.year = year
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}
