package course

import (
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"log"
	"time"
)

type (
	Service interface {
		Create(name, startDate, endDate string) (*domain.Course, error)
		Get(id string) (*domain.Course, error)
		Delete(id string) error
		Update(id string, name *string, start_date *string, end_date *string) error
		Count(filters Filters) (int, error)
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}

	Filters struct {
		Name      string
		StartDate string
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*domain.Course, error) {

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	course := &domain.Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(course); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return course, nil
}

//Obtener curso

func (s service) Get(id string) (*domain.Course, error) {
	course, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Update(id string, name *string, start_date, end_date *string) error {
	var startDateParsed, endDateParsed *time.Time

	if start_date != nil {
		date, err := time.Parse("2006-01-02", *start_date)
		if err != nil {
			s.log.Println(err)
			return err
		}
		startDateParsed = &date
	}

	if end_date != nil {
		date, err := time.Parse("2006-01-02", *end_date)
		if err != nil {
			s.log.Println(err)
			return err
		}
		endDateParsed = &date
	}
	return s.repo.Update(id, name, startDateParsed, endDateParsed)
}

func (s service) Count(filters Filters) (int, error) {

	return s.repo.Count(filters)
}

func (s service) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}
