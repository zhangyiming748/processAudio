package processAudio

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processAudio/convert"
	"github.com/zhangyiming748/processAudio/speedUp"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"strings"
)

const (
	AudioBook = "1.54" //等效audition的65%
)

func init() {
	logLevel := os.Getenv("LEVEL")
	//var level slog.Level
	var opt slog.HandlerOptions
	switch logLevel {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "processAudio.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer logf.Close() //如果不关闭可能造成内存泄露
	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(logger)
}

func ProcessAudio(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}
func ProcessAllAudio(root, pattern string) {
	ProcessAudio(root, pattern)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessAudio(folder, pattern)
	}
}
func SpeedUpAudio(dir, pattern string, speed string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		speedUp.Speedup(file, speed)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize(voiceAlert.Shanshan, strings.Join([]string{"complete", speed, "times"}, ""))
}

func SpeedUpAllAudio(root, pattern string, speed string) {
	SpeedUpAudio(root, pattern, speed)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		SpeedUpAudio(folder, pattern, speed)
	}
}
