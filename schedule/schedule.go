package schedule

import (
	"fmt"

	"releasebot/config"
	"releasebot/line"
	"releasebot/slack"

	"github.com/robfig/cron"
)

func New(config *config.Config, slack *slack.API, line *line.API) *cron.Cron {

	c := cron.New()
	for _, notice := range config.Notice {
		c.AddFunc(notice.Times, func() {
			fmt.Println("send")
			slack.PushMessage(notice.Message)
			line.PushMessage(notice.Message)
		})
	}

	return c
}