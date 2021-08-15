package model

// TODO 删掉
type Model struct {
	ID         int    `gorm:"primary_key" json:"id"` // 主键，根据约定不需要
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
