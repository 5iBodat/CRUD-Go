package designation

import "time"

type Mst_designation struct {
	DesignationId   int       `json:"designation_id" gorm:"primaryKey; autoIncrement;not_null"`
	DeptId          int       `json:"dept_id"`
	DesignationName string    `json:"designation_name"`
	CreatedAt       time.Time `json:"-"`
	LastUpdate      time.Time `json:"-" gorm:"default:null"`
}

func (Mst_designation) TableName() string {
	return "mst_designation"
}
