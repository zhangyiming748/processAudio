package processAudio

import (
	"github.com/zhangyiming748/processAudio/conv"
	"testing"
)

func TestLouderAudios(t *testing.T) {
	//LouderAllAudios("/Users/zen/container/useWget", "mp3")
	conv.SpeedUpAllAudios("/Users/zen/container/useWget", "ogg", "65")
}
