package bytea

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InsertWithV1(entity ByteaSampleEntity) error {
	db, err := gorm.Open("postgres", "host=localhost port=15432 user=myuser dbname=mydb sslmode=disable password=password")
	if err != nil {
		return fmt.Errorf("cannot connect to database: %w", err)
	}

	if err := db.Create(&entity).Error; err != nil {
		return err
	}

	return nil
}
