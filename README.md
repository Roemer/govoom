# govoom
Go Library to interact with devices from divoom (eg. Pixoo-64) via the HTTP API.

The official documentation for the API is here: https://docin.divoom-gz.com/web/#/5/23

## Installation

```go
go get github.com/roemer/govoom
```

## Usage

```go
devices, err := govoom.FindDevices()
if err != nil {
    // Handle error
}
client := devices[0].GetClient()
// Interact with the client now, eg:
client.SetBrightness(100)
```
