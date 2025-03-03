package entity

type MovieRepositoryInterface interface {
	Save(movie *Movie) error
	FindByID(id string) (*Movie, error)
	Update(movie *Movie) error
	Delete(id string) error
	GetTotal() (int, error)
	Get(filters map[string]string) ([]*Movie, error)
}
