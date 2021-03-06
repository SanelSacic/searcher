package search

import (
	"testing"
)

var final []Result

// BenchMarkRssRearch provides support for profiling the search.
func BenchMarkRssRearch(b *testing.B) {
	var result []Result
	var err error

	for i := 0; i < b.N; i++ {
		result, err = rssSearch("1", "trump", "nyt", "http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
		if err != nil {
			b.FailNow()
		}
	}

	final = result
}
