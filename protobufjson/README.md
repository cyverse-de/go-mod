# protobufjson

Implements the NATS Encoder interface, allowing NATS to automatically encode/decode messages with `protojson` through the use of an encoded connection.

## Examples

```go
func main() {
...
	flag.Parse()
	
	nats.RegisterEncoder("protojson", &protobufjson.Codec{})

...

	nc, err := nats.Connect(
		*natsURL,
		nats.UserCredentials(*credsPath),
		nats.RootCAs(*caCert),
		nats.ClientCert(*tlsCert, *tlsKey),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(*maxReconnects),
		nats.ReconnectWait(time.Duration(*reconnectWait)*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := nats.NewEncodedConn(nc, "protojson")
	if err != nil {
		log.Fatal(err)
	}
...

```

At that point NATS messages to/from that service are automatically encoded/decoded with protojson.
