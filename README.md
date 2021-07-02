# CPROXY CLI
A simple, single binary to run a forward proxy. HTTPS supported by default.

### Usage

Download the pre-built binary from https://github.com/ReeganExE/cproxy-cli/releases/latest

##### Start proxy
```sh
cproxy [options]

Options:
  -a, --addr string    Address

Example:
  cproxy                       # http://127.0.0.0:9090 (Default)
  cproxy --addr localhost:8080 # http://127.0.0.0:8080
  cproxy --addr localhost:0    # Random port http://127.0.0.0:<random_port>
  cproxy --addr localhost:     # Random port http://127.0.0.0:<random_port>
  cproxy --addr :0             # Random port http://<all_interfaces>:<random_port>
  cproxy --addr :8080          # http://<all_interfaces>:8080
```
##### Set proxy and use

The client can connect to the proxy via plaintext HTTP then stream TLS packets through the plaintext connection with perfect confidence that the connection to the specified server remains secure.

```sh
https_proxy=http://127.0.0.1:9090 curl ipinfo.io
https_proxy=http://127.0.0.1:9090 curl -L https://bing.com
https_proxy=http://127.0.0.1:9090 curl https://httpbin.org/post -d hello="from the other side"

# OR
export https_proxy=http://127.0.0.1:9090
curl https://httpbin.org/post -d hello="from the other side"
```
