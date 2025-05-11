package books

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Year      int    `gorm:"not null"`
	Publisher string `gorm:"not null"`
}
