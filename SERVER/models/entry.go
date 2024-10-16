package models

type Entry struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Dish        *string  `gorm:"type:varchar(225)" json:"dish"`
	Fat         *float64 `gorm:"type:float" json:"fat"`
	Ingredients *string  `gorm:"type:text" json:"ingredients"`
	Calories    *string  `gorm:"type:varchar(225)" json:"calories"`
}
