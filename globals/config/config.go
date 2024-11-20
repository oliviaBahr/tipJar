package config

import (
	"encoding/json"
	"os"
	fp "path/filepath"

	"tipJar/globals/log"

	"github.com/charmbracelet/lipgloss"
)

type Config struct {
	DBPath string

	AccentColor   lipgloss.Color `json:"accentColor"`
	InactiveColor lipgloss.Color `json:"inactiveColor"`
	TextColor     lipgloss.Color `json:"textColor"`

	// non-configurable. for dev use
	RepoDir string
}

func DefaultConfig() *Config {
	return &Config{
		DBPath: defaultDBPath(),

		AccentColor:   lipgloss.Color("205"),
		InactiveColor: lipgloss.Color("242"),
		TextColor:     lipgloss.Color("236"),

		RepoDir: fp.Dir(fp.Dir("config.go")),
	}
}

func LoadConfig() (*Config, error) {
	log.Debug("--- loading config ---")

	// open config file
	cfgFile, err := openConfigFile()
	if err != nil {
		log.Error("error opening config file", "e", err)
		return nil, err
	}

	// decode
	var config Config
	decoder := json.NewDecoder(cfgFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Error("error decoding config file", "e", err)
		return nil, err
	}
	cfgFile.Close()

	// set env path
	config.DBPath = userDBPath()

	return &config, nil
}

func SaveConfig(cfg *Config) error {
	log.Debug("saving config", "cfg", cfg)

	// open config file
	file, err := openConfigFile()
	if err != nil {
		log.Error("error opening config file", "e", err)
		return err
	}
	defer file.Close()

	// encode
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		log.Error("error encoding config file", "e", err)
		return err
	}

	return nil
}

func userConfigPath() (path string) {
	cfgPath := os.Getenv("TIPJAR_CONFIG_PATH")
	if !isJson(cfgPath) {
		log.Debug("TIPJAR_CONFIG_PATH env var not set or is not a json file, using default config path")
		cfgPath = defaultConfigPath()
	}
	return cfgPath
}

func userDBPath() (dbPath string) {
	dbPath = os.Getenv("TIPJAR_DB_DIR")
	if dbPath != "" {
		os.MkdirAll(fp.Dir(dbPath), 0755)
	}
	if !isDir(dbPath) {
		log.Debug("TIPJAR_DB_DIR env var not set or is not a directory, using default db path")
		dbPath = defaultDBPath()
	}
	return dbPath
}

func defaultConfigPath() string {
	var confDir = os.Getenv("XDG_CONFIG_HOME")
	if confDir == "" {
		log.Debug("XDG_CONFIG_HOME is not set, using default", "default", fp.Join(os.Getenv("HOME"), ".config"))
		confDir = fp.Join(os.Getenv("HOME"), ".config")
	}
	return fp.Join(confDir, "tipJar", "conf.json")
}

func defaultDBPath() string {
	var dataDir = os.Getenv("XDG_DATA_HOME")
	if dataDir == "" {
		log.Debug("XDG_DATA_HOME is not set, using default", "default", fp.Join(os.Getenv("HOME"), ".local", "share"))
		dataDir = fp.Join(os.Getenv("HOME"), ".local", "share")
	}

	// create db directory
	path := fp.Join(dataDir, "tipJar", "tipJar.db")
	err := os.MkdirAll(fp.Dir(path), 0755)
	if err != nil {
		log.Fatal("error creating directories", "e", err)
	}

	// create db file
	_, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("error creating db file", "e", err)
	}
	return path
}

func openConfigFile() (*os.File, error) {
	path := userConfigPath()
	log.Debug("opening config file", "path", path)

	// Create parent directories if they don't exist
	if err := os.MkdirAll(fp.Dir(path), 0755); err != nil {
		log.Error("failed to create parent dirs", "e", err)
		return nil, err
	}
	// Open file with create flag
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Error("failed to open config file", "e", err)
		return nil, err
	}
	return file, nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

func isJson(path string) bool {
	return fp.Ext(path) == ".json"
}
