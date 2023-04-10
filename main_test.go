package processAudio

import "testing"

func TestProcessAudio(t *testing.T) {
	dir := "/Users/zen/Music/Music/Media.localized/Music/小野猫/Unknown Album"
	pattern := "mp3"
	ProcessAudios(dir, pattern)
}
func TestProcessAllAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/YvDSJYMGBmjdKkPG"
	pattern := "mp3;m4a;flac;MP3;wma;wav"
	ProcessAllAudios(dir, pattern)
}
func TestSpeedUpAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/YvDSJYMGBmjdKkPG"
	pattern := "mp3;m4a;flac;MP3;wma;wav;aac"
	SpeedUpAudios(dir, pattern, AudioBook)
}
