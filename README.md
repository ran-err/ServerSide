# ServerSide

A Go server.

## How to Run

```shell
cd ServerSide

# Set up the Go workspace and resolve internal module dependencies correctly.
go work use $(find -name 'go.mod' -exec dirname {} \;)

# Generate a self-signed certificate
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes \
  -keyout cert/key.pem -out cert/cert.pem \
  -config configs/ip-cert.conf \
  -extensions v3_req
  
# Test the server with curl
curl https://127.0.0.1:8443 --cacert cert/cert.pem -v
```
