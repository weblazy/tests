package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/spf13/cast"
)

const addr = "localhost:9999"

var message = cast.ToString(time.Now().Unix())

func main() {
	tlsConf := &tls.Config{NextProtos: []string{"quic-echo-example"}, InsecureSkipVerify: true}
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		fmt.Println("err" + err.Error())
		return
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			buf := make([]byte, len(message))
			_, err = io.ReadFull(stream, buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Client: Got '%s'\n", buf)
		}

	}()
	for {
		message = cast.ToString(time.Now().Unix())
		fmt.Printf("Client: Sending '%s'\n", message)
		_, err = stream.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(2 * time.Second)
	}
}
