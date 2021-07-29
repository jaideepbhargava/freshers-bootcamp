package  Models

import (
	"fmt"
	"freshers-bootcamp/Day3/exercise2/Config"
)

func GetAllRecords (student *[]Student) (err error){
	if err = Config.DB.Find(student).Error; err != nil{
		return nil
	}

	return nil
}

func AddRecords (student *Student) (err error){
	if err = Config.DB.Create(student).Error; err != nil{
		return err
	}

	return nil
}
func GetRecordsByID(student *Student, id string) (err error){
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRecord( student *Student, id string) (err error){
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

func DeleteRecord(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil

}

