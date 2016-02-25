package data

type Post struct {
	Id string `json:"id"`
	Text string `json:"text"`
}

var latestPost = &Post{
	Id: "1",
	Text: "Hello World",
}

func GetLatestPost() *Post {
	return latestPost
}
