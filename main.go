package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	ssl_dir = "tls"

	ip       = "1.1.1.1"
	port     = "9094"
	topic    = "clickhouse.streaming"
	group_id = "go-kafka-client"
)

func main() {
	broker := ip + ":" + port
	fmt.Println("broker: ", broker)

	crt_file, key_file := ssl_dir+"/tls.crt", ssl_dir+"/tls.key"
	fmt.Println("crt_file: ", crt_file)
	fmt.Println("key_file: ", key_file)

	c := kafka.ReaderConfig{
		Brokers:        []string{broker},
		GroupID:        group_id,
		Topic:          topic,
		MinBytes:       1,
		MaxBytes:       10e6,
		StartOffset:    kafka.LastOffset,
		CommitInterval: time.Second,
		Dialer: &kafka.Dialer{
			Timeout:   1 * time.Second,
			DualStack: true,
			TLS: &tls.Config{
				InsecureSkipVerify: true,
				GetClientCertificate: func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
					cert, err := tls.LoadX509KeyPair(crt_file, key_file)
					return &cert, err
				},
			},
		},
	}

	reader := kafka.NewReader(c)
	ctx := context.Background()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("listening kafka...")
	fmt.Println()
	fmt.Println("topic: ", msg.Topic)
	fmt.Println("message: ")
	fmt.Println(string(msg.Value)[:200])
}
