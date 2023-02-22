package user

import (
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *domain.User) error
	GetAll(filters Filters, offset, limit int) ([]domain.User, error)
	Get(id string) (*domain.User, error)
	Delete(id string) error
	Update(id string, firstname *string, lastname *string, email *string, phone *string) error //path puntero, nulos y vacio
	Count(filters Filters) (int, error)
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

//instaancia repo
func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{log: log, db: db}
}

func (repo *repo) Create(user *domain.User) error {

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("repository")
	repo.log.Println("User id: ", user.Id)
	return nil
}

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {
	var u []domain.User
	tx := repo.db.Model(&u)
	tx = ApplyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (repo *repo) Get(id string) (*domain.User, error) {
	user := domain.User{Id: id}
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) Delete(id string) error {
	user := domain.User{Id: id}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) Update(id string, firstname *string, lastname *string, email *string, phone *string) error {
	values := make(map[string]interface{})
	if firstname != nil {
		values["firstname"] = *firstname
	}

	if lastname != nil {
		values["lastname"] = *lastname
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := repo.db.Model(&domain.User{}).Where("id=?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func ApplyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	//
	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName)) //todo minuscula
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)                  //filtrado
	}
	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}
	return tx
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.User{})
	tx = ApplyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
