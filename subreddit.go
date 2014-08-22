// Copyright 2012 Jimmy Zelinskie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reddit

import (
	"fmt"
  "errors"
)

// Subreddit represents a subreddit from reddit.com.
type Subreddit struct {
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

  Reddit      string   `json:"-"`
  Session     *Session `json:"-"`
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

// SubredditSubmissionss returns the submissions from this subreddit
func (s *Subreddit) SubredditSubmissionss(params map[string]interface{}) ([]*Submission, error) {
  return s.Hot(params)
}

// Comments returns the comments from this subreddit
func (s *Subreddit) Comments(params map[string]interface{}) ([]*Comment, error) {
  params["subreddit"] = Reddit
  return s.Session.SubredditComments(params)
}

func (*s Subreddit) Hot(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "hot"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) New(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "new"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) Rising(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "rising"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) Controversial(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "controversial"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) Top(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "top"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) Gilded(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "gilded"
  return s.Sesssion.SubredditSubmissionss(params)
}

func (*s Subreddit) Promoted(params map[string]interface{}) ([]*Submissions, error) {
  params["subreddit"] = Reddit
  params["find_by"] = "promoted"
  return s.Sesssion.SubredditSubmissionss(params)
}
