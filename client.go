package govoom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	url string
}

func NewClientFromDevice(d *Device) *Client {
	return NewClient(d.DevicePrivateIP)
}

func NewClient(deviceIp string) *Client {
	return &Client{
		url: fmt.Sprintf("http://%s:80/post", deviceIp),
	}
}

func (c *Client) do(data map[string]interface{}) (*http.Response, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url, &buf)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func (c *Client) validateNoError(resp *http.Response) error {
	var ret errorCodeResult
	err := json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return err
	}

	if ret.ErrorCode != 0 {
		return fmt.Errorf("command failed with error code: %d", ret.ErrorCode)
	}

	return nil
}
