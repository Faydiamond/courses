package user

import (
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"log"
)

type (
	service struct {
		repo Repository
		log  *log.Logger
	}
	//interfaces facilitan el uso d emanera generica
	Service interface {
		Create(firstName string, lastName string, email string, phone string) (*domain.User, error)
		GetAll(filters Filters, offset, limit int) ([]domain.User, error)
		Get(id string) (*domain.User, error)
		Delete(id string) error
		Update(id string, firstname *string, lastname *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
	}

	Filters struct {
		FirstName string
		LastName  string
	}
)

//llamo interfaz
func NewService(log *log.Logger, repo Repository) Service {
	return &service{log: log, repo: repo}
}

func (s service) Create(firstName, lastName, email, phone string) (*domain.User, error) {
	s.log.Println("Create user service")
	user := domain.User{FirstName: firstName, LastName: lastName, Email: email, Phone: phone}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {
	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) Get(id string) (*domain.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Update(id string, firstname *string, lastname *string, email *string, phone *string) error {
	return s.repo.Update(id, firstname, lastname, email, phone)
}
func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
