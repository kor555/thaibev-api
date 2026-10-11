package repo

import (
	"api/internal/entity/domain"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	db.AutoMigrate(&domain.AnswerChoice{})
	return &AnswerRepository{
		db: db,
	}
}

func (r *AnswerRepository) Save(answer *[]domain.AnswerChoice) error {
	result := r.db.Save(&answer)
	return result.Error
}

func (r *AnswerRepository) GetAnswerListbyQuestionID(id int) ([]domain.AnswerChoice, error) {
	list := []domain.AnswerChoice{}

	result := r.db.Where("question_id = ?", id).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (r *AnswerRepository) GetAnswerCorrectionbyID(id int) (bool, error) {
	var answer domain.AnswerChoice

	result := r.db.Where("id = ?", id).First(&answer)
	if result.Error != nil {
		return false, result.Error
	}
	return answer.IsCorrect, nil
}
