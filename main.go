package main

import (
	"flag"
	"github.com/net-byte/vtun/app"
	"github.com/net-byte/vtun/common/config"
	"os"
	"os/signal"
	"syscall"
)

var _version = "v1.6.5"

func main() {
	config := config.Config{}
	// device name 设备名称？
	flag.StringVar(&config.DeviceName, "dn", "", "device name")
	// -c 指定 tun 虚拟设备的 cidr 的地址(default "172.16.0.10/24")
	flag.StringVar(&config.CIDR, "c", "172.16.0.10/24", "tun interface cidr")
	// tun interface ipv6 cidr (default "fced:9999::9999/64")
	flag.StringVar(&config.CIDRv6, "c6", "fced:9999::9999/64", "tun interface ipv6 cidr")
	flag.IntVar(&config.MTU, "mtu", 1500, "tun mtu")
	flag.StringVar(&config.LocalAddr, "l", ":3000", "local address")
	flag.StringVar(&config.ServerAddr, "s", ":3001", "server address")
	flag.StringVar(&config.ServerIP, "sip", "172.16.0.1", "server ip")
	flag.StringVar(&config.ServerIPv6, "sip6", "fced:9999::1", "server ipv6")
	// dns server ip
	flag.StringVar(&config.DNSIP, "dip", "8.8.8.8", "dns server ip")
	flag.StringVar(&config.Key, "k", "freedom", "key")
	flag.StringVar(&config.Protocol, "p", "udp", "protocol udp/tls/grpc/ws/wss")
	flag.StringVar(&config.WebSocketPath, "path", "/freedom", "websocket path")
	// -S 设置为服务端模式
	flag.BoolVar(&config.ServerMode, "S", false, "server mode")
	flag.BoolVar(&config.GlobalMode, "g", false, "client global mode")
	flag.BoolVar(&config.Obfs, "obfs", false, "enable data obfuscation")
	flag.BoolVar(&config.Compress, "compress", false, "enable data compression")
	flag.IntVar(&config.Timeout, "t", 30, "dial timeout in seconds")
	// tls certificate file path
	flag.StringVar(&config.TLSCertificateFilePath, "certificate", "./certs/server.pem", "tls certificate file path")
	// tls certificate private key file path
	flag.StringVar(&config.TLSCertificateKeyFilePath, "privatekey", "./certs/server.key", "tls certificate key file path")
	// tls handshake sni
	flag.StringVar(&config.TLSSni, "sni", "", "tls handshake sni")
	// tls insecure skip verify
	flag.BoolVar(&config.TLSInsecureSkipVerify, "isv", false, "tls insecure skip verify")
	flag.Parse()

	app := &app.Vtun{Config: &config, Version: _version}
	app.InitConfig()
	go app.StartApp()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.StopApp()
}