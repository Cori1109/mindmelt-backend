package models

import "gorm.io/gorm"

type User_data struct {
	Vr_app_session string `gorm:"primary key;serial" json:"VR_APP_SESSION"`
	User_id string `json:"USER_ID"`
	First_name string `json:"FIRST_NAME"`
	Last_name string `json:"LAST_NAME"`
	User_role string `json:"USER_ROLE"`
	User_pass string `json:"USER_PASS"`
	Overall_score string `json:"OVERALL_SCORE"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User_data{})

	return err
}