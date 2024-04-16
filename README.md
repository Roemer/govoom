# govoom
Go Library to interact with devices from divoom (eg. Pixoo-64)

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
