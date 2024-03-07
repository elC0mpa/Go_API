package services

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"context"
	"fmt"
	"strings"
	"time"
)

func CreateBug(ctx context.Context, bug *models.Bug) error {
	return repository.InsertBug(ctx, bug)
}

func GetBugById(ctx context.Context, id uint32) (*models.Bug, error) {
	return repository.GetBugById(ctx, id)
}

func PopulateBug(ctx context.Context,  bug *models.Bug) (*models.Project, *models.User, error) {
	user, err := repository.GetUserById(ctx, bug.UserId)
	if user.Id == 0 {
		user = nil
	} else if err != nil {
		return nil, nil, err
	}
	project, err := repository.GetProjectById(ctx, bug.ProjectId)
	if project.Id == 0 {
		project = nil
	} else if err != nil {
		return nil, nil, err
	}
	return project, user, nil
}

func ListBugs(ctx context.Context, userId, projectId uint32, startDate, endDate *time.Time) ([]*models.Bug, error) {
	return repository.ListBugs(ctx, userId, projectId, startDate, endDate)
}

func BuildListBugsQuery(userID, projectID uint32, startDate, endDate *time.Time) string {
	fmt.Println("UserId", userID)
	fmt.Println("ProjectId", projectID)
	fmt.Println("StartDate", startDate)
	fmt.Println("EndDate", endDate)
	var conditions []string

	if userID != 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = %d", userID))
	}
	if projectID != 0 {
		conditions = append(conditions, fmt.Sprintf("project_id = %d", projectID))
	}
	if (startDate != nil) {
		conditions = append(conditions, fmt.Sprintf("creation_date >= '%s'", startDate.Format("2006-01-02")))
	}
	if (endDate != nil) {
		conditions = append(conditions, fmt.Sprintf("creation_date <= '%s'", endDate.Format("2006-01-02")))
	}

	query := "SELECT id, description, creation_date, user_id, project_id FROM bugs"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query
}
