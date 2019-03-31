# go-graphtweets

Build Twitter networks. `go-graphtweets` is the sister of the [`graphTweets`](https://github.com/JohnCoene/graphTweets) R package.

## Install

```bash
go get JohnCoene/go-graphtweets
```

## Example

`GetMentionsGraph` expects a struct of type `twitter.Status` as returned by [`go-twitter`](https://github.com/dghubble/go-twitter)

```go
// get tweets with go-twitter
config := oauth1.NewConfig("niLRMTSNzOw7Fqp1WNbdkdln8", "yFFYvgW8kJguFBtXmaR602Hyo1RDDRBeetZo7nIo5JecAAC4aA")
token := oauth1.NewToken("1571967648-u5ezlrsWplrHaXT4VMuVLpES2ZB8kx6UIRH4Obj", "admb6ncPJJOaCwh7HR4YASQdSJ2H9B9dIAL50yzmOHsvL")
httpClient := config.Client(oauth1.NoContext, token)

// Twitter client
client := twitter.NewClient(httpClient)

// Search Tweets
search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
  Query: "#golang filter:mentions",
})

g := graphtweets.GetMentionsGraph(*search)
```