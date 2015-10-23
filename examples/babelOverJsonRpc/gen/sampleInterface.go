
package gen

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from sample.babel ***

import (
)


// Manages users

type IUserService interface { 
	// Get a user given the ID
	// id:  User ID
	GetUser(id *int32) (*User, error) 

	// Add a user
	// user:  Details of the user to add
	AddUser(user *User) (*int32, error) 

	// CLears the list
	Clear() error 
}
