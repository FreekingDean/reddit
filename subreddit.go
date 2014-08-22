// Copyright 2012 Jimmy Zelinskie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reddit

import (
	"fmt"
)

// Subreddit represents a subreddit from reddit.com.
type Subreddit struct {
  comments
	Name        string  `json:"display_name"`
	Title       string  `json:"title"`
	Desc        string  `json:"description"`
	PublicDesc  string  `json:"public_description"`
	URL         string  `json:"url"`
	FullID      string  `json:"name"`
	ID          string  `json:"id"`
	HeaderImg   string  `json:"header_img"`
	DateCreated float32 `json:"created_utc"`
	NumSubs     int     `json:"subscribers"`
	IsNSFW      bool    `json:"over18"`

  Session     *Session `json:"-"`

  Hot           Query
  New           Query
  Rising        Query
  Controversial Queryable
  Top           Queryable
  Gilded        Query
  Promoted      Query
}

// String returns the string representation of a subreddit.
func (s *Subreddit) String() string {
	var subs string
	switch s.NumSubs {
	case 1:
		subs = "1 subscriber"
	default:
		subs = fmt.Sprintf("%d subscribers", s.NumSubs)
	}
	return fmt.Sprintf("%s (%s)", s.Title, subs)
}

// Submissions returns the submissions from this subreddit
func (s *Subreddit) Submissions(limit uint8, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.URL[3:len(s.URL)-1], limit, after)
}

// Comments returns the comments from this subreddit
func (s *Subreddit) Comments(limit uint8, after string, before string) ([]*Comment, error) {
  return s.Session.SubredditComments(s.URL[3:len(s.URL)-1], limit, after, before)
}
