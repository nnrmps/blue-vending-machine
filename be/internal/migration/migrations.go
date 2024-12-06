package migration

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Migrator struct {
	tx *gorm.DB
}

func NewMigrator(tx *gorm.DB) *Migrator {
	return &Migrator{
		tx: tx,
	}
}

func (m Migrator) Migrate() {
	migrationSteps := []*gormigrate.Migration{

		makeMigration("20241207_migration_script"),
	}
	tx := m.tx.Begin()
	gormMigrator := gormigrate.New(tx, gormigrate.DefaultOptions, migrationSteps)

	if err := gormMigrator.Migrate(); err != nil {
		tx.Rollback()
	}
	tx.Commit()

}

func makeMigration(fileName string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: fileName,
		Migrate: func(tx *gorm.DB) error {
			statements, err := readSqlFiles([]string{
				fmt.Sprintf("./internal/migration/%s.sql", fileName),
			})
			if err != nil {
				return err
			}

			for _, statement := range statements {
				err := tx.Exec(statement).Error
				if err != nil {
					return fmt.Errorf("migration error: %w", err)
				}
			}

			return nil
		},
	}
}
