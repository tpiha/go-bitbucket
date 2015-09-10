// Copyright 2015 The go-bitbucket AUTHORS. All rights reserved.
//
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"fmt"
)

// Create a new repository.
func (s *RepositoriesService) CreateHook(owner, repo string, hook *Hook) (*Hook, *Response, error) {
	var u string

	u = fmt.Sprintf("repositories/%s/%s/hooks/", owner, repo)

	req, err := s.client.NewRequest("POST", u, hook)
	if err != nil {
		return nil, nil, err
	}

	h := new(Hook)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

type Hook struct {
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Active      bool     `json:"active"`
	Events      []string `json:"events"`
}

type WebhookPayload struct {
	Actor struct {
		DisplayName string `json:"display_name"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
		Type     string `json:"type"`
		Username string `json:"username"`
		Uuid     string `json:"uuid"`
	} `json:"actor"`
	Push struct {
		Changes []struct {
			Closed  bool `json:"closed"`
			Commits []struct {
				Author struct {
					Raw string `json:"raw"`
				} `json:"author"`
				Hash  string `json:"hash"`
				Links struct {
					Html struct {
						Href string `json:"href"`
					} `json:"html"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
				Message string `json:"message"`
				Type    string `json:"type"`
			} `json:"commits"`
			Created bool `json:"created"`
			Forced  bool `json:"forced"`
			Links   struct {
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			New struct {
				Links struct {
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					Html struct {
						Href string `json:"href"`
					} `json:"html"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
				Name   string `json:"name"`
				Target struct {
					Author struct {
						Raw string `json:"raw"`
					} `json:"author"`
					Date  string `json:"date"`
					Hash  string `json:"hash"`
					Links struct {
						Html struct {
							Href string `json:"href"`
						} `json:"html"`
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
					Message string `json:"message"`
					Parents []struct {
						Hash  string `json:"hash"`
						Links struct {
							Html struct {
								Href string `json:"href"`
							} `json:"html"`
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
						} `json:"links"`
						Type string `json:"type"`
					} `json:"parents"`
					Type string `json:"type"`
				} `json:"target"`
				Type string `json:"type"`
			} `json:"new"`
			Old struct {
				Links struct {
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					Html struct {
						Href string `json:"href"`
					} `json:"html"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
				Name   string `json:"name"`
				Target struct {
					Author struct {
						Raw string `json:"raw"`
					} `json:"author"`
					Date  string `json:"date"`
					Hash  string `json:"hash"`
					Links struct {
						Html struct {
							Href string `json:"href"`
						} `json:"html"`
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
					Message string `json:"message"`
					Parents []struct {
						Hash  string `json:"hash"`
						Links struct {
							Html struct {
								Href string `json:"href"`
							} `json:"html"`
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
						} `json:"links"`
						Type string `json:"type"`
					} `json:"parents"`
					Type string `json:"type"`
				} `json:"target"`
				Type string `json:"type"`
			} `json:"old"`
			Truncated bool `json:"truncated"`
		} `json:"changes"`
	} `json:"push"`
	Repository struct {
		FullName  string `json:"full_name"`
		IsPrivate bool   `json:"is_private"`
		Links     struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
		Name  string `json:"name"`
		Owner struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"links"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Uuid     string `json:"uuid"`
		} `json:"owner"`
		Scm  string `json:"scm"`
		Type string `json:"type"`
		Uuid string `json:"uuid"`
	} `json:"repository"`
}
