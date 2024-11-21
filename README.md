# HTTPS Proxy API
An API to fetch HTTPS url content via a HTTPS tunnel proxy. Many thanks to Eli Bendersky's three part series on [Go and Proxy Servers](https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/) for the guide and informative content.

## Setup
Setting up involves running the HTTPS tunnel proxy and the API which fetches url content via the proxy.

**Tunnel Proxy**
```bash
> go run tunnel-proxy.go -addr "127.0.0.1:9999"
2024/11/11 11:11:11 starting proxy server on 127.0.0.1:9999
```

**Proxy API**
```bash
> go run proxy-api.go -addr "127.0.0.1:8080" -proxy "http://localhost:9999"
2024/11/11 11:11:11 Starting Proxy API on:  127.0.0.1:8080
2024/11/11 11:11:11 Proxy addr:  http://localhost:9999
```

## Usage
```bash
> curl localhost:8080/fetch?url=https://example.com
```

## Cases
I currently use this for my [Glance Dashboard](https://github.com/glanceapp/glance) that displays Reddit content. Since the Dashboard is hosted on a VPS and Reddit blocks VPS IPs, the Dashboard needs to fetch Reddit content via a HTTPS proxy [(link)](https://github.com/glanceapp/glance/blob/main/docs/configuration.md#reddit). The proxy is hosted on a Raspberry PI that the VPS can access via Tailscale.