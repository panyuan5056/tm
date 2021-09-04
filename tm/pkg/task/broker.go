package task

import (
	"encoding/json"
	"fmt"
	"tm/core"

	"tm/models"
	"tm/pkg/e"
	"tm/pkg/logging"

	"github.com/spf13/cast"
)

func Size() int64 {
	return models.Size()
}

func Pop() []TaskDetail {
	queues := models.Pop()
	details := []TaskDetail{}
	for _, queue := range queues {
		config := map[string]string{}
		if err := json.Unmarshal([]byte(queue.Content), &config); err == nil {
			config["type"] = queue.Category
			config["id"] = cast.ToString(queue.ID)
			fmt.Println("config:", config)
			details = append(details, TaskDetail{fn: call, params: config})
		} else {
			logging.Error(err.Error())
		}
	}
	return details
}

func call(config map[string]string) {
	var report interface{}
	if config["type"] == "1" {
		core.DbDense()
	} else if config["type"] == "2" {
		core.FileDense()
	}
	if content, err := json.Marshal(report); err == nil {
		models.Report(config["id"], string(content), e.GetMsg(e.SUCCESS), 3)
	} else {
		models.Report(config["id"], "", e.GetMsg(e.ERROR_FX), 4)
	}
}
