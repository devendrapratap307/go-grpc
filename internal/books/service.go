package books

type Service interface {
	GetBook(id uint) (*Book, error)
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(id uint) error
	ListBooks() ([]Book, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetBook(id uint) (*Book, error) { return s.repo.GetBook(id) }
func (s *service) CreateBook(book *Book) error    { return s.repo.CreateBook(book) }
func (s *service) UpdateBook(book *Book) error    { return s.repo.UpdateBook(book) }
func (s *service) DeleteBook(id uint) error       { return s.repo.DeleteBook(id) }
func (s *service) ListBooks() ([]Book, error)     { return s.repo.ListBooks() }
