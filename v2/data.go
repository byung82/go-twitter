package twitter

import "time"

type Data struct {
	ID               string            `json:"id"`
	Text             string            `json:"text"`
	AuthorID         string            `json:"author_id"`
	Lang             string            `json:"lang"`
	ReferencedTweets []ReferencedTweet `json:"referenced_tweets"`
	CreatedAt        time.Time         `json:"created_at"`
}
