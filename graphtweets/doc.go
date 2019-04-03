/*
Package graphtweets provides an easy API to construct Twitter graphs.
First set yourself up with a go-twitter client
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
Import the package with
	"go-graphtweets/graphtweets"
Then simply build retweets and mentions graphs with:
	mentionsGraph := graphtweets.GetMentionsGraph(*search) // build mentions graph
	retweetGraph := graphtweets.GetRetweetGraph(*search) // build retweet graph
	hashtagGraph := graphtweets.GetHashtagGraph(*search) // build hashtag graph
You can also convert the resulting graphs to a wider, more JSON-friendly format
	wide := graphtweets.ToWide(&mentionsGraph)
*/
package graphtweets
