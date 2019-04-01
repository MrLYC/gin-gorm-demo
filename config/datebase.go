package config

// Database : Database configuration
type Database struct {
	Type       string `yaml:"type"`
	Connection string `yaml:"connection" validate:"nonzero"`
}

// Init : init Database
func (h *Database) Init() {
	h.Type = "sqlite3"
	h.Connection = "demo.db"
}
