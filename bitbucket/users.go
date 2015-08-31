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

// ListEmails lists all email addresses for the authenticated user.
func (s *UsersService) ListEmails(opt *ListOptions) (*UserEmailList, *Response, error) {
	u := "user/emails"
	u, err := addOptions(u, opt)
	if err != nil {
		return &UserEmailList{}, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return &UserEmailList{}, nil, err
	}

	emails := new(UserEmailList)
	resp, err := s.client.Do(req, emails)
	if err != nil {
		return &UserEmailList{}, resp, err
	}

	return emails, resp, err
}

// Bitbucket user object
type User struct {
	Username    string `json:"username"`     // The name associated with the account.
	FirstName   string `json:"first_name"`   //   The first name associated with account.
	LastName    string `json:"last_name"`    // The last name associated with the account. For a team account, this value is always empty.
	DisplayName string `json:"display_name"` // Display name for the user
	IsStaff     bool   `json:"is_staff"`     // Indicates if this is a Staff account.
	Avatar      string `json:"avatar"`       // An avatar associated with the account.
	ResourceUri string `json:"resource_uri"` // Uri of the particular user through users API call
	IsTeam      bool   `json:"is_team"`      // Indicates if this is a Team account.
}

// Bitbucket user's email object
type UserEmail struct {
	IsPrimary   bool   `json:"is_primary"`   // The name associated with the account.
	IsConfirmed bool   `json:"is_confirmed"` // The name associated with the account.
	Type        string `json:"type"`         // The name associated with the account.
	Email       string `json:"email"`        // The name associated with the account.
}

// Bitbucket user's email list object
type UserEmailList struct {
	PageLen int         `json:"pagelen"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Emails  []UserEmail `json:"values"`
}
