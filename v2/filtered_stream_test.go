package twitter

import (
	"github.com/stretchr/testify/require"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestFilteredStream(t *testing.T) {
	var params url.Values
	ch := make(chan *Twitter)

	doneC, stopC, err := FilteredStream(os.Getenv("TWITTER_BEARER_TOKEN"),
		params,
		func(twitter *Twitter) {
			log.Println(twitter)
			ch <- twitter
		},
		func(err error) {
			log.Println(err)
		})

	require.NoError(t, err)

	go func() {
		<-ch
		stopC <- struct{}{}
	}()

	<-doneC
}
