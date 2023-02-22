package course

import (
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		Get(id string) (*domain.Course, error)
		Delete(id string) error
		Update(id string, name *string, start_date, end_date *time.Time) error
		Count(filters Filters) (int, error)
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(course *domain.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		r.log.Printf("error:  %v", err)
		return err
	}
	return nil
}

//obtener
func (r *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}
	if err := r.db.First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r repo) Delete(id string) error {
	couurse := domain.Course{ID: id}
	if err := r.db.Delete(&couurse).Error; err != nil {
		return err
	}
	return nil
}

func (r repo) Update(id string, name *string, start_date, end_date *time.Time) error {
	values := make(map[string]interface{})
	if name != nil {
		values["name"] = *name
	}
	if start_date != nil {
		values["start_date"] = *start_date
	}
	if end_date != nil {
		values["end_date"] = *end_date
	}

	if err := r.db.Model(&domain.Course{}).Where("id=?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func ApplyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	//
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name)) //todo minuscula
		tx = tx.Where("lower(name) like ?", filters.Name)                   //filtrado
	}
	if filters.StartDate != "" {
		filters.StartDate = fmt.Sprintf("%%%s%%", strings.ToLower(filters.StartDate))
		tx = tx.Where("lower(start_date) like ?", filters.StartDate)
	}
	return tx
}

func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(domain.Course{})
	tx = ApplyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course
	tx := repo.db.Model(&c)
	tx = ApplyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}
