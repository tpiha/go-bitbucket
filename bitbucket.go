// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://bitbucket.org/site/oauth2/"
)

// A Client manages communication with the Bitbucket API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.  Defaults to the public Bitbucket API. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the Bitbucket API.
	UserAgent string

	// Services used for talking to different parts of the Bitbucket API.
	Users *UsersService
}
