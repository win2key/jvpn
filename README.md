# jvpn
VPN automation for jump host



### Compiling
```bash
go build .
```
It will produce `jvpn` file.
### Usage
Run in console

```bash
VPN_LOGIN=vpnuser VPN_PASSWORD=vpnpassword VPN_SERVER=vpnserver.com JPROXY=:8081 ./jvpn
```

1. Open in browser `http://localhost:8080/vpn` to initiate a VPN connection.
   
2. Open in browser `http://localhost:8080/proxy` to start functioning like a `socks5` server.
   
