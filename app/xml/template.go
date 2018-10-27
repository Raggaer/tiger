package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"

	"text/template"
)

// CommandTemplate defines a controller template file
type CommandTemplate struct {
	tpl         *template.Template
	XMLName     xml.Name `xml:"template"`
	Title       string   `xml:"title"`
	Color       int      `xml:"color"`
	Description string   `xml:"description"`
}

// ParseTemplate parses the given template into a struct
func ParseTemplate(tpl string) (*CommandTemplate, error) {
	data, err := ioutil.ReadFile(tpl)
	if err != nil {
		return nil, err
	}

	// Convert XML data to struct
	templateCmd := &CommandTemplate{}
	if err := xml.Unmarshal(data, templateCmd); err != nil {
		return nil, err
	}

	// Create template data
	templateCmd.tpl = template.New("description").Funcs(templateFuncMap())
	if _, err := templateCmd.tpl.Parse(templateCmd.Description); err != nil {
		return nil, err
	}
	if _, err := templateCmd.tpl.New("title").Parse(templateCmd.Title); err != nil {
		return nil, err
	}
	return templateCmd, nil
}

// Execute executes the given command template
func (c *CommandTemplate) Execute(data map[string]interface{}) (*discordgo.MessageEmbed, error) {
	// Parse description template
	dataMap := map[string][]byte{}
	parseTemplates := []string{"description", "title"}
	for _, p := range parseTemplates {
		buff := bytes.Buffer{}
		if err := c.tpl.ExecuteTemplate(&buff, p, data); err != nil {
			return nil, err
		}
		dataMap[p] = buff.Bytes()
	}

	// Return embed message
	return &discordgo.MessageEmbed{
		Color:       c.Color,
		Description: string(dataMap["description"]),
		Title:       string(dataMap["title"]),
	}, nil
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"uptimeMessage":  uptimeMessage,
		"timeAgo":        timeAgo,
		"timeAgoCurrent": timeAgoCurrent,
		"sum":            sum,
		"unixToTime":     unixToTime,
	}
}

func unixToTime(u int64) time.Time {
	return time.Unix(u, 0)
}

func sum(a, extra int) int {
	return a + extra
}

func uptimeMessage(a, b time.Time) string {
	_, _, day, hour, min, sec := elapsedTime(a, b)
	msg := "%d day"
	if day == 0 || day > 1 {
		msg += "s"
	}
	msg += ", %d hour"
	if hour == 0 || hour > 1 {
		msg += "s"
	}
	msg += ", %d minute"
	if min == 0 || min > 1 {
		msg += "s"
	}
	msg += ", %d second"
	if sec == 0 || sec > 1 {
		msg += "s"
	}
	return fmt.Sprintf(msg, day, hour, min, sec)
}

func elapsedTime(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func timeAgoCurrent(a time.Time) string {
	return timeAgo(a, time.Now())
}

func timeAgo(a time.Time, b time.Time) string {
	y, m, d, h, x, s := elapsedTime(a, b)
	msg := ""

	// Render message as year
	if y > 0 {
		msg += strconv.Itoa(y)
		if y == 1 {
			msg += " year"
		} else {
			msg += " years"
		}
		return msg
	}

	// Render message as month
	if m > 0 {
		msg += strconv.Itoa(m)
		if m == 1 {
			msg += " month"
		} else {
			msg += " months"
		}
		return msg
	}

	// Render message as day
	if d > 0 {
		msg += strconv.Itoa(d)
		if d == 1 {
			msg += " day"
		} else {
			msg += " days"
		}
		return msg
	}

	// Render message as hour
	if h > 0 {
		msg += strconv.Itoa(h)
		if h == 1 {
			msg += " hour"
		} else {
			msg += " hours"
		}
		return msg
	}

	// Render message as minute
	if x > 0 {
		msg += strconv.Itoa(x)
		if x == 1 {
			msg += " minute"
		} else {
			msg += " minutes"
		}
		return msg
	}

	// Render message as second
	if s > 0 {
		msg += strconv.Itoa(s)
		if s == 1 {
			msg += " second"
		} else {
			msg += " seconds"
		}
		return msg
	}

	return ""
}
