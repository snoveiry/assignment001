// Package twitterv1 defines structs used by the twitter v1 endpoints
package twitterv1

type GetTweetResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
