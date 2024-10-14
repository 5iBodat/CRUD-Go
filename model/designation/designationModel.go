package designation

import (
	"golang-rest-api/config"
	"golang-rest-api/entities/designation"
)

func Read(name string) []designation.Mst_designation {
	var designations *[]designation.Mst_designation
	result := config.DB.Find(&designations, "designation_name LIKE ?", "%"+name+"%")

	if result.Error != nil {
		panic(result.Error)
	}

	return *designations
}

func Create(payload designation.Mst_designation) designation.Mst_designation {
	result := config.DB.Create(&payload)
	if result == nil {
		panic(result.Error)
	}
	return payload
}

func Update(id int, payload designation.Mst_designation) designation.Mst_designation {
	designation := designation.Mst_designation{}
	result := config.DB.First(&designation, id)

	if result.Error != nil {
		panic(result.Error)
	}
	result = config.DB.Model(&designation).Updates(payload)
	return designation
}

func Delete(id int) error {
	result := config.DB.Delete(designation.Mst_designation{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
