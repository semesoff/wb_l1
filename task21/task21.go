package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Интерфейс MediaPlayer определяет метод Play.
type MediaPlayer interface {
	Play(source string)
}

// Интерфейс AdvancedMediaPlayer определяет методы PlayAudio и PlayVideo.
type AdvancedMediaPlayer interface {
	PlayAudio(source string)
	PlayVideo(source string)
}

// AudioPlayer реализует интерфейс MediaPlayer.
type AudioPlayer struct{}

// Метод AudioPlayer
func (a *AudioPlayer) Play(source string) {
	fmt.Println("Playing audio file: ", source)
}

// VideoPlayer реализует интерфейс AdvancedMediaPlayer.
type VideoPlayer struct{}

func (v *VideoPlayer) PlayAudio(source string) {
	fmt.Println("Playing video file: ", source)
}

func (v *VideoPlayer) PlayVideo(source string) {
	fmt.Println("Playing video file: ", source)
}

// MediaAdapter реализует интерфейс MediaPlayer и содержит в себе объект типа AdvancedMediaPlayer.
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

// Метод MediaAdapter
func (m *MediaAdapter) Play(fileName string) {
	if len(fileName) <= 4 {
		fmt.Println("Invalid media. ", fileName, " format not supported")
		return
	}

	if fileName[len(fileName)-4:] == ".mp4" {
		m.advancedPlayer.PlayVideo(fileName)
	} else if fileName[len(fileName)-4:] == ".mp3" {
		m.advancedPlayer.PlayAudio(fileName)
	} else {
		fmt.Println("Invalid media. ", fileName, " format not supported")
	}
}

func main() {
	audioPlayer := &AudioPlayer{}
	videoPlayer := &VideoPlayer{}
	mediaAdapter := &MediaAdapter{advancedPlayer: videoPlayer}

	audioPlayer.Play("mp3")
	mediaAdapter.Play("yacht.mp4")
}
