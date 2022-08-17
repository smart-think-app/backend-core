package core_logging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/smart-think-app/backend-core/constant"
	"github.com/spf13/cast"
	"time"
)

type CoreLoggingManagement struct {
}

type ICoreLogging interface {
	SetInput(request interface{}) *CoreLogging
	SetOutput(response interface{}) *CoreLogging
	INFO() *CoreLogging
	ERROR() *CoreLogging
	WARN() *CoreLogging
	DONE()
}

type CoreLogging struct {
	startTime     time.Time         `json:"-"`
	Input         string            `json:"input"`
	Output        string            `json:"output"`
	ResponseTime  int64             `json:"response_time"`
	ResponseLevel string            `json:"response_level"`
	Level         string            `json:"level"`
	Error         string            `json:"error"`
	Name          string            `json:"name"`
	TrackId       string            `json:"track_id"`
	Keyword       map[string]string `json:"keyword"`
}

func NewCoreLog(startTime time.Time, name string, ctx context.Context) ICoreLogging {
	return &CoreLogging{
		startTime: startTime,
		Name:      name,
		TrackId:   cast.ToString(ctx.Value(constant.TrackID)),
	}
}

func (log *CoreLogging) SetInput(request interface{}) *CoreLogging {
	body, _ := json.Marshal(request)
	log.Input = string(body)
	return log
}

func (log *CoreLogging) SetOutput(response interface{}) *CoreLogging {
	body, _ := json.Marshal(response)
	log.Output = string(body)
	return log
}

func (log *CoreLogging) INFO() *CoreLogging {
	log.Level = constant.Info
	return log
}

func (log *CoreLogging) ERROR() *CoreLogging {
	log.Level = constant.Error
	return log
}

func (log *CoreLogging) WARN() *CoreLogging {
	log.Level = constant.Warn
	return log
}

func (log *CoreLogging) DONE() {
	now := time.Now()
	log.ResponseTime = now.Unix() -  log.startTime.Unix()
	body, _ := json.Marshal(log)
	fmt.Println(string(body))
}
