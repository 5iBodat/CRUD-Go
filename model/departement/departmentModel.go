package departement

import (
	"golang-rest-api/config"
	"golang-rest-api/entities/departement"
	"time"
)

func Read(name string) []departement.Mst_departement {
	var departements *[]departement.Mst_departement
	result := config.DB.Find(&departements, "departement_name LIKE ?", "%"+name+"%")

	if result.Error != nil {
		panic(result.Error)
	}

	return *departements
}

func Create(payload departement.Mst_departement) departement.Mst_departement {
	result := config.DB.Create(&payload)

	if result.Error != nil {
		panic(result.Error)
	}
	return payload
}

func Update(id int, payload departement.Mst_departement) departement.Mst_departement {
	departement := departement.Mst_departement{}
	result := config.DB.First(&departement, id)

	if result.Error != nil {
		panic(result.Error)
	}
	payload.LastUpdate = time.Now()
	result = config.DB.Model(&departement).Updates(payload)
	return departement
}

func Delete(id int) error {
	result := config.DB.Delete(departement.Mst_departement{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
