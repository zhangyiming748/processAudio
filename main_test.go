package processAudio

import "testing"

func TestProcessAudio(t *testing.T) {
	dir := "/Users/zen/Music/Music/Media.localized/Music/小野猫/Unknown Album"
	pattern := "mp3"
	ProcessAudio(dir, pattern)
}
func TestProcessAllAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/YvDSJYMGBmjdKkPG"
	pattern := "mp3;m4a;flac;MP3;wma;wav"
	ProcessAllAudio(dir, pattern)
}
func TestSpeedUpAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/YvDSJYMGBmjdKkPG"
	pattern := "mp3;m4a;flac;MP3;wma;wav;aac"
	SpeedUpAudio(dir, pattern, AudioBook)
}
