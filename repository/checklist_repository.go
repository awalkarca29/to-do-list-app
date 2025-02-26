package repository

import (
	"to-do-list-app/entity"

	"gorm.io/gorm"
)

type ChecklistRepository interface {
	FindAll() ([]entity.Checklist, error)
	FindByID(ID int) (entity.Checklist, error)
	Save(checklist entity.Checklist) (entity.Checklist, error)
	Update(checklist entity.Checklist) (entity.Checklist, error)
	Delete(checklist entity.Checklist) (entity.Checklist, error)
}

type checklistRepository struct {
	db *gorm.DB
}

func NewChecklistRepository(db *gorm.DB) *checklistRepository {
	return &checklistRepository{db}
}

func (r *checklistRepository) FindAll() ([]entity.Checklist, error) {
	var checklists []entity.Checklist

	err := r.db.Find(&checklists).Error
	if err != nil {
		return checklists, err
	}

	return checklists, nil
}

func (r *checklistRepository) FindByID(ID int) (entity.Checklist, error) {
	var checklist entity.Checklist

	err := r.db.Where("id = ?", ID).Find(&checklist).Error
	if err != nil {
		return checklist, err
	}
	return checklist, nil
}

func (r *checklistRepository) Save(checklist entity.Checklist) (entity.Checklist, error) {
	err := r.db.Create(&checklist).Error
	if err != nil {
		return checklist, err
	}

	return checklist, nil
}

func (r *checklistRepository) Update(checklist entity.Checklist) (entity.Checklist, error) {
	err := r.db.Save(&checklist).Error
	if err != nil {
		return checklist, err
	}

	return checklist, nil
}

func (r *checklistRepository) Delete(checklist entity.Checklist) (entity.Checklist, error) {
	err := r.db.Delete(&checklist).Error
	if err != nil {
		return checklist, err
	}

	return checklist, nil
}
