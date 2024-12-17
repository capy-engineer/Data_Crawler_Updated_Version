package model

type Swiftcode struct {
	Id                 string `json:"id" gorm:"column:id;"`
	BankInsitutionName string `json:"bank_institution_name" gorm:"column:bankName;"`
	City               string `json:"city" gorm:"column:bankCity;"`
	Branch             string `json:"branch" gorm:"column:bankBranch;"`
	SwiftCode          string `json:"swift_code" gorm:"column:bankSwift;"`
	UpdatedAt          string `json:"updated_at" gorm:"column:updatedAt;"`
	CreatedAt          string `json:"created_at" gorm:"column:createdAt;"`
}

type swiftcodeRegion struct {
	Id        string      `json:"id" gorm:"column:id;"`
	Region    string      `json:"region" gorm:"column:region;"`
	SwiftCode []Swiftcode `json:"swift_code" gorm:"foreignKey:Region;"`
}
