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
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
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
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		speedUp.Speedup(file, speed)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize(voiceAlert.Shanshan, strings.Join([]string{"complete", speed, "times"}, ""))
}

func SpeedUpAllAudio(root, pattern string, speed string) {
	SpeedUpAudio(root, pattern, speed)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		SpeedUpAudio(folder, pattern, speed)
	}
}
