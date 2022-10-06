# zkpass-alice
zkpass-node test client
## command
* run:
  - default: `go run app/main.go`
  - ping: `go run app/main.go ping`
* test：
  - `cd app/cmd`
  - `go test -v`
* build：
  - `cd app`
  - `go build -o ../alice`
* install：
  - `go list -f '{{.Target}}'`
  - `go env -w GOBIN=/path/to/your/bin`
  - `go install`