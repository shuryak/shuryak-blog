package main

import (
	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)
