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
func SpeedupAudio(in GetFileInfo.Info, speed string) {
	src := strings.Trim(in.FullPath, in.FullName)   //原文件目录 带有最后一个 /
	dst := strings.Join([]string{src, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	outName := strings.Trim(in.FullName, in.ExtName)
	outName = strings.Join([]string{outName, "aac"}, "")
	slog.Debug("补全后的 outName", slog.String("outName", outName))
	target := strings.Join([]string{dst, outName}, string(os.PathSeparator))
	slog.Info("io", slog.String("输入文件", in.FullPath), slog.String("输出文件", target))
	sppedUp(in.FullPath, target, speed)
}

func sppedUp(in, out string, speed string) {
	atempo := strings.Join([]string{"atempo", speed}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vn", out)
	slog.Info("生成命令", slog.String("命令", fmt.Sprint(cmd)))
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
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("错误", err))
		return
	}
	if err = os.RemoveAll(in); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in), slog.Any("错误内容", err))
	} else {
		slog.Info("删除成功", slog.String("源文件", in))
	}
}
