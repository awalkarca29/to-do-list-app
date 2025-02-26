package helper

import "to-do-list-app/entity"

type ChecklistFormatter struct {
	ID     int    `json:"id"`
	Title  string `json:"description"`
	UserID int    `json:"user_id"`
}

func FormatChecklist(checklist entity.Checklist) ChecklistFormatter {
	checklistFormatter := ChecklistFormatter{}
	checklistFormatter.ID = checklist.ID
	checklistFormatter.Title = checklist.Title
	checklistFormatter.UserID = checklist.UserID

	return checklistFormatter
}

func FormatChecklists(checklists []entity.Checklist) []ChecklistFormatter {
	checklistsFormatter := []ChecklistFormatter{}

	for _, checklist := range checklists {
		checklistFormatter := FormatChecklist(checklist)
		checklistsFormatter = append(checklistsFormatter, checklistFormatter)
	}

	return checklistsFormatter
}

type ChecklistDetailFormatter struct {
	ID     int    `json:"id"`
	Title  string `json:"name"`
	UserID int    `json:"user_id"`
}

func FormatChecklistDetail(checklist entity.Checklist) ChecklistDetailFormatter {
	checklistDetailFormatter := ChecklistDetailFormatter{}
	checklistDetailFormatter.ID = checklist.ID
	checklistDetailFormatter.Title = checklist.Title
	checklistDetailFormatter.UserID = checklist.UserID

	return checklistDetailFormatter
}
