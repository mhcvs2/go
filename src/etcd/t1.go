package main

import (
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/pkg/transport"
	"os"
	"time"
)

func t11() {
	os.Setenv("ETCDCTL_CA_FILE", "/etc/etcd/ssl/etcd-ca.pem")
	os.Setenv("ETCDCTL_CERT_FILE", "/etc/etcd/ssl/etcd.pem")
	os.Setenv("ETCDCTL_KEY_FILE", "/etc/etcd/ssl/etcd-key.pem")

	endpoints := []string{"https://109.105.1.253:2379"}

	cafile := os.Getenv("ETCDCTL_CA_FILE")
	certfile := os.Getenv("ETCDCTL_CERT_FILE")
	keyfile := os.Getenv("ETCDCTL_KEY_FILE")

	tls := transport.TLSInfo{
		CertFile: certfile,
		KeyFile: keyfile,
		TrustedCAFile: cafile,
	}
	dialTimeout := 5 * time.Second

	cfg := client.Config{
		Endpoints: endpoints,
		Transport: transport.NewTransport(tls, dialTimeout),
	}
}
