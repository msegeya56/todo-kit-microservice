package connection

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)






var DB *gorm.DB


















type Arguments struct {
	Host     string
	Port     int
	Secure   bool
	SSLMode  string
	Username string
	Password string
	TargetDB string
}







func connection(args *Arguments) (*gorm.DB, error) {
	if args == nil {
		DB, err := gorm.Open(sqlite.Open("Customers.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return DB, nil
	} else {
		if len(args.TargetDB) == 0 {
			args.TargetDB = "postgres"
		}

		if len(args.Host) == 0 {
			return nil, errors.New("Missing Arguments.Host")
		}

		if len(args.Username) == 0 {
			return nil, errors.New("Missing Arguments.Username")
		}

		if len(args.Password) == 0 {
			return nil, errors.New("Missing Arguments.Password")
		}

		if args.Port == 0 {
			args.Port = 5432
		}

		if args.Port == 0 {
			return nil, errors.New("Missing Arguments.Port")
		}

		if args.Secure {
			args.SSLMode = "require"
		} else {
			args.SSLMode = "disabled"
		}

		switch args.TargetDB {
		case "postgres":
			// Handle PostgreSQL connection
		}
	}

	return DB, nil
}