package speedUp

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

func Speedup(in GetFileInfo.Info, speed string) {
	src := strings.Trim(in.FullPath, in.FullName)   //原文件目录 带有最后一个 /
	dst := strings.Join([]string{src, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	target := strings.Join([]string{dst, in.FullName}, string(os.PathSeparator))
	slog.Info("io", slog.Any("输入文件", in.FullPath), slog.Any("输出文件", target))
	sppedUp_help(in.FullPath, target, speed)
}

func sppedUp_help(in, out string, speed string) {
	atempo := strings.Join([]string{"atempo", speed}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vn", out)
	slog.Info("command", slog.Any("生成的命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("产生的错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("产生的错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		fmt.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("命令执行中", slog.Any("产生的错误", err))
		return
	}
	//if err := os.RemoveAll(in); err != nil {
	//			slog.Warn("删除失败", slog.Any("源文件", in.FullPath))
	//} else {
	//	slog.Warn("删除成功", slog.Any("源文件", in.FullPath))
	//}
}
