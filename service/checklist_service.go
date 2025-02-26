package service

import (
	"errors"
	"to-do-list-app/entity"
	"to-do-list-app/repository"
)

type GetChecklistDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateChecklistInput struct {
	Title string `form:"title" binding:"required"`
	User  entity.User
}

type ChecklistService interface {
	GetAllChecklists() ([]entity.Checklist, error)
	GetChecklistByID(input GetChecklistDetailInput) (entity.Checklist, error)
	CreateChecklist(input CreateChecklistInput) (entity.Checklist, error)
	DeleteChecklist(inputID GetChecklistDetailInput) (entity.Checklist, error)
}

type checklistService struct {
	checklistRepository repository.ChecklistRepository
}

func NewChecklistService(checklistRepository repository.ChecklistRepository) *checklistService {
	return &checklistService{checklistRepository}
}

func (s *checklistService) GetAllChecklists() ([]entity.Checklist, error) {
	checklists, err := s.checklistRepository.FindAll()
	if err != nil {
		return checklists, err
	}
	return checklists, nil
}

func (s *checklistService) GetChecklistByID(input GetChecklistDetailInput) (entity.Checklist, error) {
	checklist, err := s.checklistRepository.FindByID(input.ID)

	if err != nil {
		return checklist, err
	}

	if checklist.ID == 0 {
		return checklist, errors.New("no checklist found on that id")
	}

	return checklist, nil
}

func (s *checklistService) CreateChecklist(input CreateChecklistInput) (entity.Checklist, error) {
	checklist := entity.Checklist{}
	checklist.Title = input.Title
	checklist.UserID = input.User.ID

	newChecklist, err := s.checklistRepository.Save(checklist)
	if err != nil {
		return newChecklist, err
	}

	return newChecklist, nil
}

func (s *checklistService) DeleteChecklist(inputID GetChecklistDetailInput) (entity.Checklist, error) {
	checklist, err := s.checklistRepository.FindByID(inputID.ID)
	if err != nil {
		return checklist, err
	}

	if checklist.ID == 0 {
		return checklist, errors.New("no checklist found on that id")
	}

	deleteChecklist, err := s.checklistRepository.Delete(checklist)
	if err != nil {
		return deleteChecklist, err
	}

	return deleteChecklist, nil
}
