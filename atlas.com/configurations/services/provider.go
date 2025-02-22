package services

import (
	"atlas-configurations/database"
	"context"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func allEntityProvider(ctx context.Context) database.EntityProvider[[]Entity] {
	return func(db *gorm.DB) model.Provider[[]Entity] {
		var results []Entity
		err := db.WithContext(ctx).Find(&results).Error
		if err != nil {
			return model.ErrorProvider[[]Entity](err)
		}
		return model.FixedProvider[[]Entity](results)
	}
}

func byIdEntityProvider(ctx context.Context) func(id uuid.UUID) database.EntityProvider[Entity] {
	return func(id uuid.UUID) database.EntityProvider[Entity] {
		return func(db *gorm.DB) model.Provider[Entity] {
			var result Entity
			err := db.WithContext(ctx).Where("id = ?", id).First(&result).Error
			if err != nil {
				return model.ErrorProvider[Entity](err)
			}
			return model.FixedProvider[Entity](result)
		}
	}
}
