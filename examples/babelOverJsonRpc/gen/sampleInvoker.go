
package gen

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from sample.babel ***

import (
)


// UserServiceGetUserRequest is the request structure used for invoking the GetUser method on the UserService service.
type UserServiceGetUserRequest struct {

	// User ID
	Id *int32 `json:"id,omitempty"`

}

// Init sets default values for a GetUserRequest
func (obj *UserServiceGetUserRequest) Init() *UserServiceGetUserRequest {
	return obj
}

// UserServiceGetUserResponse is the response structure used for invoking the GetUser method on the UserService service.
type UserServiceGetUserResponse struct {

	Value *User `json:"Value,omitempty"`

}

// Init sets default values for a GetUserResponse
func (obj *UserServiceGetUserResponse) Init() *UserServiceGetUserResponse {
	return obj
}


// UserServiceAddUserRequest is the request structure used for invoking the AddUser method on the UserService service.
type UserServiceAddUserRequest struct {

	// Details of the user to add
	User *User `json:"user,omitempty"`

}

// Init sets default values for a AddUserRequest
func (obj *UserServiceAddUserRequest) Init() *UserServiceAddUserRequest {
	return obj
}

// UserServiceAddUserResponse is the response structure used for invoking the AddUser method on the UserService service.
type UserServiceAddUserResponse struct {

	Value *int32 `json:"Value,omitempty"`

}

// Init sets default values for a AddUserResponse
func (obj *UserServiceAddUserResponse) Init() *UserServiceAddUserResponse {
	return obj
}


// UserServiceClearRequest is the request structure used for invoking the Clear method on the UserService service.
type UserServiceClearRequest struct {

}

// Init sets default values for a ClearRequest
func (obj *UserServiceClearRequest) Init() *UserServiceClearRequest {
	return obj
}

// UserServiceClearResponse is the response structure used for invoking the Clear method on the UserService service.
type UserServiceClearResponse struct {

}

// Init sets default values for a ClearResponse
func (obj *UserServiceClearResponse) Init() *UserServiceClearResponse {
	return obj
}




// Manages users
type UserService struct {
	SvcObj IUserService `json:"-"`
}

// Get a user given the ID
func (s *UserService) GetUser(req *UserServiceGetUserRequest, rsp *UserServiceGetUserResponse) error {
	response, err := s.SvcObj.GetUser(req.Id)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Add a user
func (s *UserService) AddUser(req *UserServiceAddUserRequest, rsp *UserServiceAddUserResponse) error {
	response, err := s.SvcObj.AddUser(req.User)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// CLears the list
func (s *UserService) Clear(req *UserServiceClearRequest, rsp *UserServiceClearResponse) error {
	err := s.SvcObj.Clear()

	return err
}



