package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	AudioBook = "1.54" //等效audition的65%
)

var (
	mylog *slog.Logger
)

func setLog(level string) {
	var opt slog.HandlerOptions
	switch level {
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
	mylog = slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
}
func init() {
	level := os.Getenv("LEVEL")
	setLog(level)
}
func processAudio(in GetFileInfo.Info) {
	out := strings.Join([]string{strings.Trim(in.FullPath, in.ExtName), "aac"}, "")
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, out)
	mylog.Debug("生成命令", slog.String("命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe", slog.Any("错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run", slog.Any("错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err := cmd.Wait(); err != nil {
		mylog.Warn("命令执行中", slog.Any("错误", err))
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		mylog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		mylog.Debug("删除成功", slog.String("源文件", in.FullPath))
	}
}

func ProcessAudios(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		processAudio(file)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}
func ProcessAllAudios(root, pattern string) {
	ProcessAudios(root, pattern)
	folders := GetAllFolder.List(root)
	for _, folder := range folders {
		ProcessAudios(folder, pattern)
	}
}
func SpeedUpAudios(dir, pattern string, speed string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		SpeedupAudio(file, speed)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize(voiceAlert.Samantha, strings.Join([]string{"complete", speed, "times"}, ""))
}

func SpeedUpAllAudios(root, pattern string, speed string) {
	SpeedUpAudios(root, pattern, speed)
	folders := GetAllFolder.List(root)
	for _, folder := range folders {
		SpeedUpAudios(folder, pattern, speed)
	}
}
