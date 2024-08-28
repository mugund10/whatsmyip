> [!NOTE]
> The [server](https://aws.mugund10.top/ip) is secured with HTTPS using TLS certificates obtained through my [LetsEncryptAcmeClient](https://github.com/mugund10/LetsEncryptAcmeClient)


## whatsmyip

A [simple HTTP server](https://aws.mugund10.top/ip) that echoes a client public IP address. This is utilized in my [Dynamic DNS (DDNS)](https://github.com/mugund10/dynamicdns) project to verify the current public IP and ensure accurate DNS updates.

## usage

```bash
$ curl https://aws.mugund10.top/ip
```
`{"addr":"120.56.34.0"}`
