package twitter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type FilteredStreamHandler func(twitter *Twitter)

func FilteredStream(bearerToken string,
	params url.Values,
	handler FilteredStreamHandler,
	errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {

	endPoint := fmt.Sprintf("%s/tweets/search/stream?%s", apiUrl,
		params.Encode())

	cfg := NewConfig(bearerToken, endPoint)

	recvHandler := func(recvBytes []byte) {
		if len(recvBytes) > 2 {
			var twitter Twitter

			if err := json.Unmarshal(recvBytes, &twitter); err == nil {
				handler(&twitter)
			} else {
				log.Println(err)
			}
		}
	}
	return streamServer(cfg, recvHandler, errHandler)
}
