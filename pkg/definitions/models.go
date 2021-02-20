package definitions

import (
	"time"
)

// Config ...
type Config struct {
	Migrations   string   `yaml:"migrations"`
	Models       []string `yaml:"models"`
	AddGormModel bool     `yaml:"add_gorm_model"`
}

// Model is a same model as defines in `gorm.Model`
// for gorm-migrations, models need not specify `gorm.Model`
// can't figure out why types.Check can't find it's import path
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// FieldMeta ...
type FieldMeta struct {
	Tag       string
	FieldType BasicType
}