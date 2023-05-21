package processAudio

import "testing"

func TestProcessAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/Telegram"
	pattern := "mp3;aac"
	ConvAudios(dir, pattern)
}
func TestProcessAllAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/YvDSJYMGBmjdKkPG"
	pattern := "mp3;m4a;flac;MP3;wma;wav"
	ConvAllAudios(dir, pattern)
}
func TestSpeedUpAudio(t *testing.T) {
	dir := "/Users/zen/Downloads/Telegram"
	pattern := "mp3;m4a;flac;MP3;wma;wav;aac"
	SpeedUpAudios(dir, pattern, AudioBook)
}
func TestSpeedAndConv(t *testing.T) {
	setLog("Debug")
	dir := "/Users/zen/github/ConvAudio/storage"
	pattern := "mp3;m4a;flac;MP3;wma;wav;aac"
	SpeedUpAudios(dir, pattern, AudioBook)
}
