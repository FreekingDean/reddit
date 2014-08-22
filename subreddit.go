// Copyright 2012 Jimmy Zelinskie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reddit

import (
	"fmt"
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

// Submissions returns the submissions from this subreddit
func (s *Subreddit) Submissions(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Hot(period, limit, before, after)
}

// Comments returns the comments from this subreddit
func (s *Subreddit) Comments(limit uint8, before string, after string) ([]*Comment, error) {
  return s.Session.SubredditComments(s.Reddit, limit, before, after)
}

// Hot returns the subreddits submissions sorted by "hot"
func (s *Subreddit) Hot(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "hot", period, limit, before, after)
}

// New returns the subreddits submissions sorted by "new"
func (s *Subreddit) Not(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "new", period, limit, before, after)
}

// Rising returns the subreddits submissions sorted by "rising"
func (s *Subreddit) Rising(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "rising", period, limit, before, after)
}

// Controversial returns the subreddits submissions sorted by "controversial"
func (s *Subreddit) Controversial(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "controversial", period, limit, before, after)
}

// Top returns the subreddits submissions sorted by "top"
func (s *Subreddit) Top(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "top", period, limit, before, after)
}

// Gilded returns the subreddits submissions sorted by "gilded"
func (s *Subreddit) Gilded(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "gilded", period, limit, before, after)
}

// Promoted returns the subreddits submissions sorted by "promoted"
func (s *Subreddit) Promoted(period string, limit uint8, before string, after string) ([]*Submission, error) {
  return s.Session.SubredditSubmissions(s.Reddit, "promoted", period, limit, before, after)
}
