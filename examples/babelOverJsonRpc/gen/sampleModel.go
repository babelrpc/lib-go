
package gen

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from sample.babel ***

import (
)




// User data
type User struct {
	// Name of the user
	Name *string `json:"name,omitempty"`

	// Age of the user
	Age *int32 `json:"age,omitempty"`

	// Email address of the user
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Internally assigned user id
	Id *int32 `json:"id,omitempty"`

}

// Init sets default values for a User
func (obj *User) Init() *User {
	return obj
}
