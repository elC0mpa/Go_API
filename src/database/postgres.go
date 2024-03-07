package database

import (
	"api-rest/src/models"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
	
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (name, surname) VALUES ($1, $2)", user.Name, user.Surname)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id uint32) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, surname FROM users WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Surname); err == nil {
			return &user, nil
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (repo *PostgresRepository) InsertProject(ctx context.Context, project *models.Project) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO projects (name, description) VALUES ($1, $2)", project.Name, project.Description)
	return err
}

func (repo *PostgresRepository) GetProjectById(ctx context.Context, id uint32) (*models.Project, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, description FROM projects WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var project = models.Project{}
	for rows.Next() {
		if err = rows.Scan(&project.Id, &project.Name, &project.Description); err == nil {
			return &project, nil
		} else {
			return nil, err
		}
	}
	return &project, nil
}

func (repo *PostgresRepository) InsertBug(ctx context.Context, bug *models.Bug) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO bugs (description, user_id, project_id) VALUES ($1, $2, $3)", bug.Description, bug.UserId, bug.ProjectId)
	return err
}

func (repo *PostgresRepository) GetBugById(ctx context.Context, id uint32) (*models.Bug, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, description, creation_date, user_id, project_id FROM bugs WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var bug = models.Bug{}
	for rows.Next() {
		if err = rows.Scan(&bug.Id, &bug.Description, &bug.CreationDate, &bug.UserId, &bug.ProjectId); err == nil {
			return &bug, nil
		} else {
			return nil, err
		}
	}
	return &bug, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
