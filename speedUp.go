package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

func SpeedupAudio(in GetFileInfo.Info, speed string) {
	src := strings.Trim(in.FullPath, in.FullName)   //原文件目录 带有最后一个 /
	dst := strings.Join([]string{src, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	target := strings.Join([]string{dst, in.FullName}, string(os.PathSeparator))
	mylog.Info("io", slog.Any("输入文件", in.FullPath), slog.Any("输出文件", target))
	sppedUp(in.FullPath, target, speed)
}

func sppedUp(in, out string, speed string) {
	atempo := strings.Join([]string{"atempo", speed}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vn", out)
	mylog.Info("生成命令", slog.String("命令", fmt.Sprint(cmd)))
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
	//if err := os.RemoveAll(in); err != nil {
	//			mylog.Warn("删除失败", slog.Any("源文件", in.FullPath))
	//} else {
	//	mylog.Info("删除成功", slog.Any("源文件", in.FullPath))
	//}
}
