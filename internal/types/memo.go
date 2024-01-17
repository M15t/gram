package types

// Memo represents the memo model
// swagger:model
type Memo struct {
	Base
	UserID  string `json:"-"  gorm:"not null"`
	Content string `json:"memo" gorm:"type:text"`
}
