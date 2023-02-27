package processAudio

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processAudio/convert"
	"github.com/zhangyiming748/processAudio/speedUp"
	"github.com/zhangyiming748/voiceAlert"
	"strings"
)

const (
	AudioBook = "1.54" //等效audition的65%
)

func ProcessAudio(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "任务执行失败")
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个任务转换成功")
	}
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "整个任务执行完成")
}
func ProcessAllAudio(root, pattern string) {
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessAudio(folder, pattern)
	}
}
func SpeedUpAudio(dir, pattern string, speed string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "加速任务失败")
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		speedUp.Speedup(file, speed)
		voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个音频加速成功")
	}
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, strings.Join([]string{"全部音频加速", speed, "倍成功"}, ""))
}
func SpeedUpAllAudio(root, pattern string, speed string) {
	SpeedUpAudio(root, pattern, speed)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		SpeedUpAudio(folder, pattern, speed)
	}
}
