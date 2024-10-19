package postgres_counter

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CounterTable struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Count     int
}

func (CounterTable) TableName() string {
	return "_counter"
}

func Counter(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&CounterTable{})

	for i := 0; ; i++ {
		record := CounterTable{Count: i}

		if err := db.Create(&record).Error; err != nil {
			fmt.Println("Error inserting record:", err)
			return
		}

		fmt.Printf("Inserted record: ID = %d, Count = %d, CreatedAt = %s\n", record.ID, i, record.CreatedAt.Format(time.RFC3339))

		time.Sleep(1 * time.Second)
	}
}
