package speedUp

import (
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

func Speedup(in GetFileInfo.Info, speed string) {
	src := strings.Trim(in.FullPath, in.FullName)   //原文件目录 带有最后一个 /
	dst := strings.Join([]string{src, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	log.Debug.Printf("src = %v\tdst = %v\n", src, dst)
	target := strings.Join([]string{dst, in.FullName}, string(os.PathSeparator))
	log.Debug.Printf("target = %v\n", target)
	sppedUp_help(in.FullPath, target, speed)
}

func sppedUp_help(in, out string, speed string) {
	atempo := strings.Join([]string{"atempo", speed}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vn", out)
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
	//if err := os.RemoveAll(in); err != nil {
	//	log.Warn.Printf("删除源文件失败:%v\n", err)
	//} else {
	//	log.Debug.Printf("删除源文件:%v\n", in)
	//}
}
