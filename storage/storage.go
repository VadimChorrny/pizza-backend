package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"pizza-backend/storage/models"
	"time"
)

type Storage struct {
	Db *gorm.DB
}

type Config struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// New creates new storage instance.
func New(url string, config Config) (*Storage, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}
	return &Storage{
		Db: db,
	}, nil
}

// DB gets db.
func (s *Storage) DB() *gorm.DB {
	return s.Db
}

func (s *Storage) Begin() *Storage {
	return &Storage{
		Db: s.Db.Begin(),
	}
}

func (s *Storage) Commit() error {
	return s.Db.Commit().Error
}

func (s *Storage) Rollback() error {
	return s.Db.Rollback().Error
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

func NotFound() error {
	return gorm.ErrRecordNotFound
}

func (s *Storage) CreateUser(user *models.User) error {
	return s.Db.Create(&user).Error
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.Db.Where("email = ?", email).First(&user).Error
	return &user, err
}
