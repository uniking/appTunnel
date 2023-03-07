#!/bin/bash
go run proxyClient.go -src 127.0.0.1:3182
go run tunnel.go -src 0.0.0.0:8081 -dest 127.0.0.1:3182 -crypt -server