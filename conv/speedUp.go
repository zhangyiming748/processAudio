package conv

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processAudio/replace"
	"github.com/zhangyiming748/processAudio/util"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	//AudioBook = "1.54" //等效audition的65%
	AudioBook = "1.43" //等效audition的70%

)

/*
执行文件夹和子文件夹中音频加速
*/
func SpeedUpAllAudios(root, pattern string, speed string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		SpeedupAudio(in, speed)
	}
}

/*
执行一个文件夹中音频加速
*/
func SpeedUpAudios(dir, pattern string, speed string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		SpeedupAudio(in, speed)
	}
}

/*
加速单个音频的完整函数
*/
func SpeedupAudio(in GetFileInfo.BasicInfo, speed string) {
	dst := strings.Join([]string{in.PurgePath, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	fname := replace.ForFileName(in.PurgeName)
	fname = strings.Join([]string{fname, "ogg"}, ".")
	slog.Debug("补全后的 fname", slog.String("fname", fname))
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
	//跳过已经加速的文件夹
	if strings.Contains(in.FullPath, "speed") {
		return
	}
	speedUp(in.FullPath, out, speed)
}

/*
仅使用输入输出和加速参数执行命令
*/
func speedUp(in, out string, speed string) {
	ff := audition2ffmpeg(speed)
	atempo := strings.Join([]string{"atempo", ff}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vn", "-ac", "1", "-map_metadata", "-1", out)
	util.ExecCommand(cmd)
	//if err := os.RemoveAll(in); err != nil {
	//	slog.Warn("删除失败", slog.String("源文件", in), slog.Any("错误内容", err))
	//} else {
	//	slog.Debug("删除成功", slog.String("源文件", in))
	//}
}

/*
获取一个等效adobe audition 的 混缩
*/
func audition2ffmpeg(speed string) string {
	audition, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		slog.Warn("解析加速参数错误,退出程序", slog.String("错误原文", fmt.Sprint(err)))
		os.Exit(1)
	}
	param := 100 / audition
	slog.Debug("转换后的原始参数", slog.Float64("param", param))
	final := fmt.Sprintf("%.2f", param)
	slog.Debug("保留两位小数的原始参数", slog.String("final", final))
	return final
}
