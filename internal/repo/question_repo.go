package repo

import (
	"api/internal/entity/domain"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	db.AutoMigrate(&domain.Question{})
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) Save(question *domain.Question) (*domain.Question, error) {
	result := r.db.Save(&question)
	return question, result.Error
}

func (r *QuestionRepository) GetQuestionbyID(id int) (*domain.Question, error) {
	var questions domain.Question

	result := r.db.Where("id = ?", id).Find(&questions)
	if result.Error != nil {
		return &domain.Question{}, result.Error
	}
	return &questions, nil
}

func (r *QuestionRepository) GetAllQuestionID() (*[]domain.Question, error) {
	var idList []domain.Question

	result := r.db.Select("ID").Find(&idList)
	if result.Error != nil {
		return &[]domain.Question{}, result.Error
	}
	return &idList, nil
}
