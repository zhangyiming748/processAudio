package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

const (
	AudioBook = "1.54" //等效audition的65%
)

func ConvAudio(in GetFileInfo.Info) {
	out := strings.Join([]string{strings.Trim(in.FullPath, in.ExtName), "aac"}, "")
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, out)
	slog.Debug("生成命令", slog.String("命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err = stdout.Read(tmp)
		t := string(tmp)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("错误", err))
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err = os.RemoveAll(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		slog.Debug("删除成功", slog.String("源文件", in.FullPath))
	}
}

func ConvAudios(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		ConvAudio(file)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}
func ConvAllAudios(root, pattern string) {
	ConvAudios(root, pattern)
	folders := GetAllFolder.List(root)
	for _, folder := range folders {
		ConvAudios(folder, pattern)
	}
}
