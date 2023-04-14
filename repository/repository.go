package repository

import (
	"github.com/sam-explorex/demo_math_grpc/entities"
	"gorm.io/gorm"
)

// create log for any operation func CreateEmails(email *entities.Emails) (entities.Emails, error) {
func CreateLogs(log entities.Logs, Db *gorm.DB) (entities.Logs, error) {
	err := Db.Create(&log).Error
	if err != nil {
		return entities.Logs{}, err

	}
	return log, nil
}
