package models

import (
	"time"
)

// Product ...
type Product struct {
	Code  string
	Price uint
}

// User ...
type User struct {
	Name string
	Age  uint
}

// // IsModel ..
// func (m *User) IsModel() bool {
// 	return true
// }

// Organisations is the parent of each individual account
type Organisations struct {
	GUID             *string `gorm:"not null;unique;column:guid"`
	Name             *string `gorm:"not null;column:name"`
	OrganisationCode *int32  `gorm:"not null;unique;column:organisation_code"`
}

// Accounts is the child/children of an organisation. An organisation must have at least one
// acoount
type Accounts struct {
	GUID              *string   `gorm:"not null;unique;column:guid"`
	FirstName         *string   `gorm:"not null;column:first_name"`
	LastName          *string   `gorm:"not null;column:last_name"`
	OrganisationID    *string   `gorm:"not null;column:organisation_id"`
	IsAccountBillable *bool     `gorm:"default:true;column:is_account_billable"`
	APIKey            *string   `gorm:"not null;type:text;column:api_key"`
	Type              *string   `gorm:"not null;type:varchar(255);column:type"`
	Active            *bool     `gorm:"default:true"`
	HasAcceptedTerms  bool      `gorm:"default:false"`
	AmountPaid        float32   `gorm:"null;type:numeric"`
	AmountDeducted    *float32  `gorm:"type:numeric;default:0"`
	Date              time.Time `gorm:"not null"`
	// Grouped           pq.StringArray `gorm:"type:varchar(64)[];"`
}

// IsModel ..
func (m *Accounts) IsModel() bool {
	return true
}

// Credentials holds auth credentials when user logs in via afya notes console
// This table is populated on first sign up
type Credentials struct {
	GUID      string `gorm:"not null;unique;column:guid"`
	AccountID string `gorm:"not null;unique;column:account_id"`
	Password  string `gorm:"type:varchar(255)"`
}