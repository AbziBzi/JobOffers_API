package enums

// Role describes user role
type Role struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}
