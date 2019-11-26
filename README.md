[![Go Report Card](https://goreportcard.com/badge/github.com/JohnCoene/go-graphtweets)](https://goreportcard.com/report/github.com/JohnCoene/go-graphtweets) [![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/JohnCoene/go-graphtweets/graphtweets)

# go-graphtweets

Build Twitter networks. 

_`go-graphtweets` is the sister of the [`graphTweets`](https://github.com/JohnCoene/graphTweets) R package._

## Install

Install `go-twitter`

```bash
go get github.com/dghubble/go-twitter/twitter
```

Install `go-graphtweets`

```bash
go get github.com/JohnCoene/go-graphtweets/graphtweets
```

## Documentation

The documentation is on [godoc](https://godoc.org/github.com/JohnCoene/go-graphtweets/graphtweets)

## Example

`GetMentionsGraph` expects a struct of type `twitter.Status` as returned by [`go-twitter`](https://github.com/dghubble/go-twitter).

Set up `go-twitter` session.

```go
// get tweets with go-twitter
config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
httpClient := config.Client(oauth1.NoContext, token)

// Twitter client
client := twitter.NewClient(httpClient)

// Search Tweets
search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
  Query: "#golang",
})
```

Build graphs.

```go
mentionsGraph := graphtweets.GetMentionsGraph(*search) // build mentions graph
retweetGraph := graphtweets.GetRetweetGraph(*search) // build retweet graph
hashtagGraph := graphtweets.GetHashtagGraph(*search) // build hashtag graph
```

Convert either graphs to wide.

```go
wide := graphtweets.ToWide(&mentionsGraph)
```

Wide format is generally better/closer to the typical `JSON`.

```go
serialized,_ := json.Marshal(wide)

fmt.Println(string(serialized))
```