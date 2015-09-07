// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"fmt"
)

type ChangesetsService struct {
	client *Client
}

// list changesets
func (cs *ChangesetsService) List(owner, repo string, limit int, start string) (*ChangesetList, *Response, error) {
	var u string

	u = fmt.Sprintf("repositories/%s/%s/changesets", owner, repo)

	req, err := cs.client.NewRequestV1("POST", u, repo)
	if err != nil {
		return nil, nil, err
	}

	csl := new(ChangesetList)
	resp, err := cs.client.Do(req, csl)
	if err != nil {
		return nil, resp, err
	}

	return csl, resp, err
}

type Changeset struct {
	Node  string `json:"node"`
	Files []struct {
		Type string `json:"type"`
		File string `json:"file"`
	} `json:"files"`
	Branches     []interface{} `json:"branches"`
	RawAuthor    string        `json:"raw_author"`
	Utctimestamp string        `json:"utctimestamp"`
	Author       string        `json:"author"`
	Timestamp    string        `json:"timestamp"`
	RawNode      string        `json:"raw_node"`
	Parents      []string      `json:"parents"`
	Branch       interface{}   `json:"branch"`
	Message      string        `json:"message"`
	Revision     interface{}   `json:"revision"`
	Size         int           `json:"size"`
}

type ChangesetList struct {
	Count      int          `json:"count"`
	Start      string       `json:"start"`
	Limit      int          `json:"limit"`
	Changesets []*Changeset `json:"changesets"`
}
