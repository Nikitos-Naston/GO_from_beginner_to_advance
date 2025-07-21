package storage

import (
	"fmt"
	"learning_GO/RAMEN_API_base/internal/app/models"
	"log"
)

// Должен Добавлять / удалять по id , получать все статьи , получать статьи по id, обновлять статьи по id
type ArcticleRepository struct {
	storage *Storage
}

var (
	tableArcticle string = "articles"
)

func (ar *ArcticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArcticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *ArcticleRepository) DeleteById(id int) (*models.Article, error) {
	arcticle, ok, err := ar.FindArcicleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArcticle)
		_, err := ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return arcticle, nil
}

func (ar *ArcticleRepository) FindArcicleById(id int) (*models.Article, bool, error) {
	arcticals, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var arcticalFind *models.Article
	for _, a := range arcticals {
		if a.ID == id {
			arcticalFind = a
			founded = true
			break
		}
	}
	return arcticalFind, founded, nil
}

// получим всех пользователей из бд
func (ar *ArcticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArcticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arcticals := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		arcticals = append(arcticals, &a)
	}
	return arcticals, nil
}
