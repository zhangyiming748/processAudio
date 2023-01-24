package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processAudio/convert"
	"github.com/zhangyiming748/processAudio/util"
	"os"
	"strings"
	"time"
)

func ProcessAudio(dir, pattern string) {
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)

	var files []util.File
	files = util.GetMultiFiles(dir, pattern)
	for _, file := range files {
		convert.Convert2AAC(file)
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}
func ProcessAllAudio(root, pattern string) {
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	var files []util.File
	folders := listFolders(root)
	for _, src := range folders {
		files = util.GetMultiFiles(src, pattern)
		for _, file := range files {
			convert.Convert2AAC(file)
		}
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}

func listFolders(dirname string) []string {
	fileInfos, _ := os.ReadDir(dirname)
	var folders []string
	for _, fi := range fileInfos {
		filename := strings.Join([]string{dirname, fi.Name()}, "/") //拼写当前文件夹中所有的文件地址
		//fmt.Println(filename)                                       //打印文件地址
		if fi.IsDir() { //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(fi.Name(), "h265") {
				continue
			}
			fmt.Printf("获取到的文件夹:%v\n", filename)
			folders = append(folders, filename)
			listFolders(filename) //递归调用
		}
	}
	return folders
}
