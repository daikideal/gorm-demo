package bytea

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InsertWithV2(entity ByteaSampleEntity) error {
	db, err := gorm.Open(postgres.Open("host=localhost port=15432 user=myuser dbname=mydb sslmode=disable password=password"))
	if err != nil {
		return fmt.Errorf("cannot connect to database: %w", err)
	}

	if err := db.Create(&entity).Error; err != nil {
		return err
	}

	return nil
}
