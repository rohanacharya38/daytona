// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package db

import (
	. "github.com/daytonaio/daytona/pkg/db/dto"
	"github.com/daytonaio/daytona/pkg/profiledata"
	"gorm.io/gorm"
)

type ProfileDataStore struct {
	db *gorm.DB
}

func NewProfileDataStore(db *gorm.DB) (*ProfileDataStore, error) {
	err := db.AutoMigrate(&ProfileDataDTO{})
	if err != nil {
		return nil, err
	}

	return &ProfileDataStore{db: db}, nil
}

func (s *ProfileDataStore) Get() (*profiledata.ProfileData, error) {
	profileDataDTO := ProfileDataDTO{}
	tx := s.db.Where("id = ?", ProfileDataId).First(&profileDataDTO)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, profiledata.ErrProfileDataNotFound
		}
		return nil, tx.Error
	}

	profileData := ToProfileData(profileDataDTO)

	return profileData, nil
}

func (s *ProfileDataStore) Save(profileData *profiledata.ProfileData) error {
	profileDataDTO := ToProfileDataDTO(profileData)
	tx := s.db.Save(&profileDataDTO)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (s *ProfileDataStore) Delete() error {
	tx := s.db.Where("id = ?", ProfileDataId).Delete(&ProfileDataDTO{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return profiledata.ErrProfileDataNotFound
	}

	return nil
}
