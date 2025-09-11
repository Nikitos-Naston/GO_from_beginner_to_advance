package models

import "Ginogorm/storage"

func GetallArticles(a *[]Arcticle) error {
	if err := storage.DB.Find(a).Error; err != nil {
		return err
	}
	return nil
}

func AddNewArticle(a *Arcticle) error {
	if err := storage.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func GetArticlebyID(a *Arcticle, id string) error {
	if err := storage.DB.Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticlebyID(a *Arcticle, id string) error {
	if err := storage.DB.Where("id = ?", id).Delete(a).Error; err != nil {
		return err
	}
	return nil
}

func UpdateArticlebyID(a *Arcticle, id string) error {
	if err := storage.DB.Save(a).Error; err != nil {
		return err
	}
	return nil
}
