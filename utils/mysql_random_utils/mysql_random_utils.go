package mysql_random_utils

import (
	"strconv"

	"github.com/sikalabs/slu/utils/random_utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Example struct {
	Foo string
	Bar string
	Baz string
}

func GenerateExample() Example {
	return Example{
		Foo: random_utils.RandomString(32, random_utils.LOWER),
		Bar: random_utils.RandomString(32, random_utils.LOWER),
		Baz: random_utils.RandomString(32, random_utils.LOWER),
	}
}

func GenerateRandomData(
	host string,
	port int,
	user string,
	password string,
	name string,
	batchSize int,
	batchCount int,
) {
	db, err := gorm.Open(
		mysql.Open(user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+name),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Example{})
	var examples = make([]Example, batchSize)

	for i := range examples {
		examples[i] = GenerateExample()
	}

	for i := 1; i < batchCount; i++ {
		db.Create(&examples)
	}
}
