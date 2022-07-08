package utils

import "github.com/sirupsen/logrus"

// 打印日志时，多输出一个字段
type LogTrace struct {
}

func NewLogTrace() LogTrace {
	return LogTrace{}
}

func (hook LogTrace) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LogTrace) Fire(entry *logrus.Entry) error {
	ctx := entry.Context
	if ctx != nil {
		traceId := ctx.Value("trace_id")
		if traceId != nil {
			entry.Data["trace_id"] = traceId
		}
	}
	return nil
}
