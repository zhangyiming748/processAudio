package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

func Convert2AAC(in GetFileInfo.Info) {
	out := strings.Join([]string{strings.Trim(in.FullPath, in.ExtName), "aac"}, ".")
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, out)
	mylog.Debug("生成命令", slog.String("命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe", slog.Any("产生的错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run", slog.Any("产生的错误", err))
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
	if err := cmd.Wait(); err != nil {
		mylog.Warn("命令执行中", slog.Any("产生的错误", err))
		return
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		mylog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		mylog.Warn("删除成功", slog.String("源文件", in.FullPath))
	}
}
