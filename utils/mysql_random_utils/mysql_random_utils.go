package mysql_random_utils

import (
	"strconv"

	"github.com/sikalabs/slu/utils/random_utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Example struct {
	Alpha   string
	Bravo   string
	Charlie string
	Delta   string
	Echo    string
}

func GenerateExample() Example {
	return Example{
		Alpha:   random_utils.RandomString(64, random_utils.LOWER),
		Bravo:   random_utils.RandomString(64, random_utils.LOWER),
		Charlie: random_utils.RandomString(64, random_utils.LOWER),
		Delta:   random_utils.RandomString(64, random_utils.LOWER),
		Echo:    random_utils.RandomString(64, random_utils.LOWER),
	}
}

func GenerateRandomData(
	host string,
	port int,
	user string,
	password string,
	databaseName string,
	tableName string,
	batchSize int,
	batchCount int,
) {
	if tableName == "" {
		tableName = "examples"
	}

	db, err := gorm.Open(
		mysql.Open(user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+databaseName),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic(err)
	}

	db.Table(tableName).AutoMigrate(&Example{})
	var examples = make([]Example, batchSize)

	for i := range examples {
		examples[i] = GenerateExample()
	}

	for i := 1; i < batchCount; i++ {
		db.Table(tableName).Create(&examples)
	}
}
