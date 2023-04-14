package repository

import (
	"github.com/sam-explorex/demo_math_grpc/entities"
	"gorm.io/gorm"
)

type LogRepository struct {
	Db *gorm.DB
}

// create log for any operation
func (l LogRepository) CreateLogs(log entities.Logs) (entities.Logs, error) {
	err := l.Db.Create(&log).Error
	if err != nil {
		return entities.Logs{}, err

	}
	return log, nil
}
