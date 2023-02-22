package enrollment

import (
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"log"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
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

func (r *repo) Create(enroll *domain.Enrollment) error {

	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("enrollment creado con el siguiente id: ", enroll.ID)
	return nil
}
