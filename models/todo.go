package models

type Todo struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Description string `gorm:"type:varchar(300)" json:"description"`
	CreatedAt   int    `gorm:"autoCreateTime:mili" json:"createdAt"`
	ClosedAt    int    `gorm:"autoUpdateTime:mili" json:"closedAt"`
	Completed   int    `gorm:"default:0" json:"completed"`
}
