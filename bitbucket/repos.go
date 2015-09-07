// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"fmt"
)

// Repository represents a Bitbucket repository.
type Repository struct {
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	Owner    string      `json:"owner"`
	Scm      string      `json:"scm"`
	Logo     string      `json:"logo"`
	Language string      `json:"language"`
	Private  bool        `json:"is_private"`
	IsFork   bool        `json:"is_fork"`
	ForkOf   *Repository `json:"fork_of"`
}

// RepositoriesService handles communication with the repository related
// methods of the Bitbucket API.
type RepositoriesService struct {
	client *Client
}

// Create a new repository.
func (s *RepositoriesService) Create(repo *Repository) (*Repository, *Response, error) {
	var u string

	u = fmt.Sprintf("repositories/%s/%s", repo.Owner, repo.Slug)

	req, err := s.client.NewRequest("POST", u, repo)
	if err != nil {
		return nil, nil, err
	}

	r := new(Repository)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}
