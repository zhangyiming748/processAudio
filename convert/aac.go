package convert

import (
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

func Convert2AAC(in GetFileInfo.Info) {
	out := strings.Join([]string{strings.Trim(in.FullPath, in.ExtName), "aac"}, ".")
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, out)
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.TTY.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
}
