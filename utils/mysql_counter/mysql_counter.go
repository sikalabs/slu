package mysql_counter

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
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
	db, err := gorm.Open(mysql.Open(user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+name), &gorm.Config{})
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
