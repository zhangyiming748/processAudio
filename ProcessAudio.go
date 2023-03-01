package processAudio

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processAudio/convert"
	"github.com/zhangyiming748/voiceAlert"
)

func ProcessAudio(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	if len(files) == 0 {
		voiceAlert.Customize("skip", voiceAlert.Samantha)
	}
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllAudio(root, pattern string) {
	ProcessAudio(root, pattern)
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessAudio(folder, pattern)
	}
}
