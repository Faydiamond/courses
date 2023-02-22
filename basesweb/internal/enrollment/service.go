package enrollment

import (
	"curso_golang/Faydiamond/basesweb/internal/course"
	"curso_golang/Faydiamond/basesweb/internal/domain"
	"curso_golang/Faydiamond/basesweb/internal/user"
	"errors"
	"log"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}
	service struct {
		log       *log.Logger
		userSrv   user.Service
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(l *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:       l,
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		//me interesa el error pq no existe
		return nil, errors.New("Usuario no existe ")
	}

	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil {
		return nil, errors.New("curso no existe")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
