package govoom

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func (c *Client) ResetSendingAnimationPicId() error {
	cmd := "Draw/ResetHttpGifId"
	data := map[string]interface{}{
		"Command": cmd,
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}

type getSendingAnimationPicIDResult struct {
	errorCodeResult
	PicId int `json:"PicId"`
}

func (c *Client) GetSendingAnimationPicID() (int, error) {
	cmd := "Draw/GetHttpGifId"
	data := map[string]interface{}{
		"Command": cmd,
	}

	resp, err := c.do(data)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	var ret getSendingAnimationPicIDResult
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return -1, err
	}

	if ret.ErrorCode != 0 {
		return -1, fmt.Errorf("GetSendingAnimationPicID failed: %d", ret.ErrorCode)
	}

	return ret.PicId, nil
}

func (c *Client) SendAnimation(animationId int, numPictures int, pictureIndex int, width int, speed int, picData []byte) error {
	cmd := "Draw/SendHttpGif"
	data := map[string]interface{}{
		"Command":   cmd,
		"PicNum":    numPictures,
		"PicWidth":  width,
		"PicOffset": pictureIndex,
		"PicID":     animationId,
		"PicSpeed":  speed,
		"PicData":   base64.StdEncoding.EncodeToString(picData),
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}
