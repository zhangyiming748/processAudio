package processAudio

import (
	"fmt"
	"github.com/zhangyiming748/processAudio/conv"
	"os"
)

// libmp3lame
func main() {
	root := ""
	if len(os.Args) > 1 {
		root = os.Args[1]
		fmt.Println("请输入要转换的文件夹")
	}
	pattern := "mp3;m4a;flac;wma;wav;m4a;aac;ogg"
	conv.AllAudios2Mp3(root, pattern)

}
