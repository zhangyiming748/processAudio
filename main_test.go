package processAudio

import "testing"

func TestProcessAudio(t *testing.T) {
	dir := "/Users/zen/Music/Music/Media.localized/Music/小野猫/Unknown Album"
	pattern := "mp3"
	ProcessAudio(dir, pattern)
}
func TestProcessAllAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/Affect3d/Audio Updates"
	pattern := "mp3;wav"
	ProcessAllAudio(dir, pattern)
}
