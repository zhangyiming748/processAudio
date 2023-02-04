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
			voiceAlert.Voice(voiceAlert.FAILED)
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
		voiceAlert.Voice(voiceAlert.SUCCESS)
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
	voiceAlert.Voice(voiceAlert.COMPLETE)
}
func ProcessAllAudio(root, pattern string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Voice(voiceAlert.FAILED)
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	ProcessAudio(root, pattern)
	folders := GetAllFolder.ListFolders(root)
	for _, dir := range folders {
		files := GetFileInfo.GetAllFileInfo(dir, pattern)
		for _, file := range files {
			convert.Convert2AAC(file)
			voiceAlert.Voice(voiceAlert.SUCCESS)
		}
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	voiceAlert.Voice(voiceAlert.COMPLETE)
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}
