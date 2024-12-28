package core

import (
	"database/sql"
	"os"
	fp "path/filepath"
	"tipJar/globals/config"

	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

type Jar struct {
	*sql.DB
	Tips []*Tip
	Cfg  *config.Config
}

func LoadJar(cfg *config.Config) (*Jar, error) {
	log.Debug("--- loading jar ---")

	// open db
	log.Debug("opening db", "dbPath", cfg.DBPath)
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Error("error opening db", "e", err)
		return nil, err
	}

	// create table
	log.Debug("creating table")
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tipJar (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		tags TEXT,
		links TEXT);
	`)
	if err != nil {
		log.Error("error opening or creating table", "e", err)
		db.Close()
		return nil, err
	}

	log.Debug("loading tips from db")
	tips, err := db.Query("SELECT * FROM tipJar")
	if err != nil {
		log.Error("error querying database", "e", err)
		db.Close()
		return nil, err
	}
	defer tips.Close()

	jar := &Jar{db, rowsToTips(tips), cfg}
	return jar, nil
}

func LoadTestJar() (*Jar, error) {
	log.Debug("--- loading test jar ---")
	cfg := config.DefaultConfig()

	// Create temp directory that will be removed on program exit
	log.Debug("creating temp directory")
	tempDir, err := os.MkdirTemp("", "tipjar-test-*")
	if err != nil {
		log.Error("error creating temp directory", "e", err)
		return nil, err
	}

	log.Debug("loading jar", "dbPath", fp.Join(tempDir, "tipJar.db"))
	cfg.DBPath = fp.Join(tempDir, "tipJar.db")
	jar, err := LoadJar(cfg)
	if err != nil {
		log.Error("error loading test jar", "e", err)
		return nil, err
	}

	log.Debug("deleting all tips from jar")
	jar.Exec("DELETE FROM tipJar")
	return jar, nil
}

func (j *Jar) Query(query string, args ...any) []*Tip {
	rows, err := j.DB.Query(query, args...)
	if err != nil {
		log.Error("error querying database", "e", err)
		return nil
	}
	defer rows.Close()
	return rowsToTips(rows)
}

func rowsToTips(rows *sql.Rows) []*Tip {
	tips := []*Tip{}
	for rows.Next() {
		var id, title, description, tags, links string
		if err := rows.Scan(&id, &title, &description, &tags, &links); err != nil {
			log.Error("error scanning rows", "e", err)
			return nil
		}
		tip := NewOldTip(id, title, description, tags, links)
		tips = append(tips, tip)
	}
	return tips
}
