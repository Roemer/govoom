package govoom

type TextScrollDirection int

const (
	TextScrollDirectionLeft  TextScrollDirection = 0
	TextScrollDirectionRight TextScrollDirection = 1
)

type TextFont int

const (
	TextFont0 TextFont = iota
	TextFont1
	TextFont2
	TextFont3
	TextFont4
	TextFont5
	TextFont6
	TextFont7
)

type TextAlignment int

const (
	TextAlignmentLeft   TextAlignment = 1
	TextAlignmentMiddle TextAlignment = 2
	TextAlignmentRight  TextAlignment = 3
)

func (c *Client) ClearAllTextArea() error {
	cmd := "Draw/ClearHttpText"
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

func (c *Client) SendText(id, x, y int, dir TextScrollDirection, font TextFont, width int, str string, speed int, color string, align TextAlignment) error {
	cmd := "Draw/SendHttpText"
	data := map[string]interface{}{
		"Command":    cmd,
		"TextId":     id,
		"x":          x,
		"y":          y,
		"dir":        int(dir),
		"font":       int(font),
		"TextWidth":  width,
		"TextString": str,
		"speed":      speed,
		"color":      color,
		"align":      int(align),
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}

type TextType int

const (
	TextTypeTimeSecond           TextType = 1
	TextTypeTimeMinute           TextType = 2
	TextTypeTimeHour             TextType = 3
	TextTypeTimeAmPm             TextType = 4
	TextTypeTimeHourMinute       TextType = 5 // hour:min
	TextTypeTimeHourMinuteSecond TextType = 6 // hour:min:sec
	TextTypeDateYear             TextType = 7
	TextTypeDateDay              TextType = 8
	TextTypeDateMonth            TextType = 9
	TextTypeDateMonthYear        TextType = 10 // month-year
	TextTypeDateMonthNameDotDay  TextType = 11 // month-name.day
	TextTypeDateWeekYear         TextType = 12 // day:month:year ???
	TextTypeDateDayName2Char     TextType = 13 // weekday-"SU","MO","TU","WE","TH","FR","SA"
	TextTypeDateDayName3Char     TextType = 14 // weekday-"SUN","MON","TUE","WED","THU","FRI","SAT"
	TextTypeDateDayNameFull      TextType = 15 // weekday-"SUNDAY","MONDAY","TUESDAY","WEDNESDAY","THURSDAY","FRIDAY","SATURDAY"
	TextTypeDateMonthName3Char   TextType = 16 // month-"JAN","FEB","MAR","APR","MAY","JUN","JUL","AUG","SEP","OCT","NOV","DEC"
	TextTypeWeatherTemp          TextType = 17 // temperature + c/f
	TextTypeWeatherTodayMaxTemp  TextType = 18
	TextTypeWeatherTodayMinTemp  TextType = 19
	TextTypeWeatherName          TextType = 20 // the weather
	TextTypeNoise                TextType = 21 // the noise value
	TextTypeText                 TextType = 22 // a defined text
	TextTypeNetText              TextType = 23 // an url that responds with a json including the "DispData" string element, eg:http://appin.divoom-gz.com/Device/ReturnCurrentDate?test=0 reponse {"DispData": "2022-01-22 13:51:56"}
)

type DisplayListElement struct {
	Id              int                 `json:"TextId"`
	TextType        TextType            `json:"type"`
	X               int                 `json:"x"`
	Y               int                 `json:"y"`
	ScrollDirection TextScrollDirection `json:"dir"`
	Font            int                 `json:"font"`
	Width           int                 `json:"TextWidth"`
	Height          int                 `json:"Textheight"`
	Text            string              `json:"TextString"`
	Speed           int                 `json:"speed"`
	Color           string              `json:"color"`
	UpdateTime      int                 `json:"update_time"`
	TextAlignment   TextAlignment       `json:"align"`
}

func (c *Client) SendDisplayList(elements ...DisplayListElement) error {
	cmd := "Draw/SendHttpItemList"
	data := map[string]interface{}{
		"Command":  cmd,
		"ItemList": elements,
	}

	resp, err := c.do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return c.validateNoError(resp)
}
