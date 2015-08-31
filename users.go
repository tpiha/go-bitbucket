// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

// UsersService handles communication with the user related
// methods of the Bitbucket API.
type UsersService struct {
	client *Client
}

// Returns current authenticated user object
func (s *UsersService) Current() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user", nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

// Bitbucket user object
type User struct {
	Username    string `json:"username"`     // The name associated with the account.
	FirstName   string `json:"first_name"`   //   The first name associated with account.
	LastName    string `json:"last_name"`    // The last name associated with the account. For a team account, this value is always empty.
	DisplayName string `json:"display_name"` // Display name for the user
	IsStaff     bool   `json:"is_staff"`     // An avatar associated with the account.
	Avatar      string `json:"avatar"`       // An avatar associated with the account.
	ResourceUri string `json:"resource_uri"` // Uri of the particular user through users API call
	IsTeam      bool   `json:"is_team"`      // Indicates if this is a Team account.
}
