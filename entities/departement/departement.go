package departement

import "time"

type Mst_departement struct {
	DepartementId       int       `json:"departement_id" gorm:"primaryKey; autoIncrement;not_null"`
	DepartementName     string    `json:"departement_name"`
	DepartementParentId int       `json:"departement_parent_id"`
	DepartementPicId    int       `json:"departement_pic_id"`
	CreatedAt           time.Time `json:"-"`
	LastUpdate          time.Time `json:"-" gorm:"default:null"`
	IsDeleted           bool      `json:"-" gorm:"default:false"`
}

func (Mst_departement) TableName() string {
	return "mst_departement"
}
