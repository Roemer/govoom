package govoom

import (
	"fmt"
)

var (
	ErrInvalidBrightness = fmt.Errorf("brightness should be in range of 0-100")
)

func (c *Client) SetBrightness(brightness int) error {
	if brightness < 0 || brightness > 100 {
		return ErrInvalidBrightness
	}

	cmd := "Channel/SetBrightness"
	data := map[string]interface{}{
		"Command":    cmd,
		"Brightness": brightness,
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}

type RotationAngle int

const (
	RotationNormal RotationAngle = 0
	Rotation90     RotationAngle = 1
	Rotation180    RotationAngle = 2
	Rotation270    RotationAngle = 3
)

func (c *Client) SetRotation(deviceIp string, rotation RotationAngle) error {
	cmd := "Device/SetScreenRotationAngle"
	data := map[string]interface{}{
		"Command": cmd,
		"Mode":    int(rotation),
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}

type MirrorMode int

const (
	MirrorModeDisable MirrorMode = 0
	MirrorModeEnable  MirrorMode = 1
)

func (c *Client) SetMirrorMode(mirrorMode MirrorMode) error {
	cmd := "Device/SetMirrorMode"
	data := map[string]interface{}{
		"Command": cmd,
		"Mode":    int(mirrorMode),
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}
