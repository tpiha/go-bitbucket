// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"bytes"
	"fmt"
	"io"
)

// RepositoriesServiceV1 handles communication with the repository related
// methods of the Bitbucket API 1.0.
type RepositoriesServiceV1 struct {
	client *Client
}

// Gets content of a file
func (s *RepositoriesServiceV1) GetContent(owner, repo, sha, path string) (*FileContentData, *Response, error) {
	var u string

	u = fmt.Sprintf("repositories/%s/%s/src/%s/%s", owner, repo, sha, path)

	req, err := s.client.NewRequestV1("GET", u, repo)
	if err != nil {
		return nil, nil, err
	}

	fc := new(FileContent)
	resp, err := s.client.Do(req, fc)
	if err != nil {
		return nil, resp, err
	}

	fcd := &FileContentData{bytes.NewBufferString(fc.Data)}

	return fcd, resp, err
}

type FileContent struct {
	Node string `json:"node"`
	Path string `json:"path"`
	Data string `json:"data"`
	Size int    `json:"size"`
}

type FileContentData struct {
	io.Reader
}

func (fcd *FileContentData) Close() error {
	return nil
}
