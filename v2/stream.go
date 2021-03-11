package twitter

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
)

// RecvHandler handle raw twitter message
type RecvHandler func(recvBytes []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

var streamServer = func(cfg *Config, handler RecvHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	c := http.Client{
		Transport: &http2.Transport{},
	}

	fmt.Println(cfg.EndPoint)

	req, _ := http.NewRequest("GET",
		cfg.EndPoint,
		nil)
	req.Header.Add("Authorization", "Bearer "+cfg.BearerToken) // 헤더값 설정

	resp, err := c.Do(req)

	if err != nil {
		log.Fatalf("Failed get: %s", err)
		return nil, nil, err
	}

	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)

		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			_ = resp.Body.Close()
		}()

		bufferedReader := bufio.NewReader(resp.Body)

		log.Printf("Status: %s\n", resp.Status)

		if resp.StatusCode >= 400 {
			errHandler(errors.New(resp.Status))
			stopC <- struct{}{}
			return
		}

		for {
			recvBytes, err := bufferedReader.ReadString('\n')

			if len(recvBytes) > 10 {
				log.Printf("recvBytes:%d %s\n", len(recvBytes), recvBytes)
			}

			if err != nil {
				if err != io.EOF {
					if !silent {
						errHandler(err)
					}
				}

				return
			}

			if len(recvBytes) > 10 {
				handler([]byte(recvBytes))
			}

		}
	}()
	return
}
