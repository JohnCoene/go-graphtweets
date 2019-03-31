# go-graphtweets

Build Twitter networks. `go-graphtweets` is the sister of the [`graphTweets`](https://github.com/JohnCoene/graphTweets) R package.

## Install

Install `go-twitter`

```bash
go get JohnCoene/go-graphtweets
```

Install `go-graphtweets`

```bash
go get dghubble/go-twitter
```

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

Build graphs

```go
mentionsGraph := graphtweets.GetMentionsGraph(*search) // build graph
rtweetGraph := graphtweets.GetRetweetGraph(*search) // build graph
```