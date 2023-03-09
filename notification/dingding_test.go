package notification_test

import (
	"testing"

	"github.com/chaimch/akita-inu/notification"
	"github.com/sirupsen/logrus"
)

func TestSendDD(t *testing.T) {
	ddMsg := &notification.DDMsg{
		Msgtype: notification.MsgTypeMD,
		Markdown: notification.DDMsgMarkDown{
			Title: "Warning: gitlab project name conflict!",
			Text:  "### xxx",
		},
	}

	robot := "https://oapi.dingtalk.com/robot/send?access_token=xxxx"

	resp, err := notification.SendDD(ddMsg, robot)
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Info(resp.Errcode)
	logrus.Info(resp.Errmsg)

}
