package processAudio

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processAudio/convert"
	"github.com/zhangyiming748/voiceAlert"
	"time"
)

func ProcessAudio(dir, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "任务执行失败")
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个任务转换成功")
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "整个任务执行完成")
}
func ProcessAllAudio(root, pattern string) {
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		ProcessAudio(folder, pattern)
	}
}
