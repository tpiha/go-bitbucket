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
	Name        string      `json:"name,omitempty"`
	Slug        string      `json:"slug,omitempty"`
	Owner       string      `json:"owner,omitempty"`
	Scm         string      `json:"scm,omitempty"`
	Logo        string      `json:"logo,omitempty"`
	Language    string      `json:"language,omitempty"`
	Description string      `json:"description,omitempty"`
	Private     bool        `json:"is_private,omitempty"`
	IsFork      bool        `json:"is_fork,omitempty"`
	ForkOf      *Repository `json:"fork_of,omitempty"`
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

// Delete a repository.
func (s *RepositoriesService) Delete(owner, repo string) (*Response, error) {
	var u string

	u = fmt.Sprintf("repositories/%s/%s/", owner, repo)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
