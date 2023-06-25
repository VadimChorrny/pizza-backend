package migrations

import (
	"gorm.io/gorm"
	"pizza-backend/storage"
	"pizza-backend/utils"
)

type Version struct {
	Version int `gorm:"primary_key"`
}

func Migrate(db *storage.Storage) error {
	err := db.DB().AutoMigrate(&Version{})
	if err != nil {
		return err
	}

	ver := Version{}
	rs := db.DB().Model(&Version{}).Order("version desc").First(&ver)
	if rs.Error != nil && gorm.ErrRecordNotFound != rs.Error {
		return err
	}

	for index, v := range migrations {
		if index+1 > ver.Version { // skip old migrations
			tx := db.DB().Begin()
			err := tx.Create(&Version{Version: index + 1}).Error
			if err != nil {
				return err
			}

			if err = tx.Exec(v).Error; err != nil {
				return err
			}

			err = tx.Commit().Error
			if err != nil {
				return err
			}
		}
	}

	utils.Logger().Info().Msg("migrations applied")

	return nil
}
