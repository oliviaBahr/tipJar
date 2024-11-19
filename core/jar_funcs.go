package core

import (
	"fmt"
	"strings"
)

func (j *Jar) AddTip(title, description, tags, links string) error {
	tip := NewTip(title, description, tags, links)
	_, err := j.Exec("INSERT INTO tipJar (id, title, description, tags, links) VALUES (?, ?, ?, ?, ?)", tip.ID, tip.Title, tip.Description, strings.Join(tip.Tags, ","), strings.Join(tip.Links, ","))
	if err != nil {
		return err
	}
	return nil
}

func (j *Jar) RemoveTip(tip *Tip) error {
	if tip.ID == "" {
		panic("tip ID is empty")
	}

	fmt.Printf("Removing tip with ID: %s\n", tip.ID)

	existing := j.Query("SELECT * FROM tipJar WHERE id = ?", tip.ID)
	if len(existing) == 0 {
		return fmt.Errorf("tip with ID %s not found in database", tip.ID)
	}

	res, err := j.Exec("DELETE FROM tipJar WHERE id = ?", tip.ID)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}
	aff, _ := res.RowsAffected()
	if aff < 1 {
		return fmt.Errorf("tip with ID %s not found", tip.ID)
	}

	return nil
}

func (j *Jar) GetAllTips() []*Tip {
	return j.Query("SELECT * FROM tipJar")
}

func (j *Jar) SearchTips(query string, tags []string) (tips []*Tip) {
	hasQuery := query != ""
	hasTags := len(tags) > 0

	switch {
	case hasQuery && hasTags:
		tips = j.Query("SELECT * FROM tipJar WHERE tags LIKE ? AND (title LIKE ? OR description LIKE ? OR links LIKE ?)", "%"+strings.Join(tags, ",")+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	case !hasQuery && hasTags:
		tips = j.Query("SELECT * FROM tipJar WHERE tags LIKE ?", "%"+strings.Join(tags, ",")+"%")
	case hasQuery && !hasTags:
		tips = j.Query("SELECT * FROM tipJar WHERE title LIKE ? OR description LIKE ? OR links LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	default:
		tips = j.Query("SELECT * FROM tipJar")
	}

	return tips
}

func (j *Jar) SearchByTags(tags []string) []*Tip {
	return j.Query("SELECT * FROM tipJar WHERE tags LIKE ?", "%"+strings.Join(tags, ",")+"%")
}
