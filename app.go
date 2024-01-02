package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
)

//go:generate stringer -type=FType
type FType uint8

const (
	image FType = iota
	video
	audio
)

type FileStats struct {
	FileType FType
	Regex    *regexp.Regexp
	Files    []fs.DirEntry
}

func AppendIfMatch(stat *FileStats, file fs.DirEntry) {
	if stat.Regex.MatchString(file.Name()) {
		stat.Files = append(stat.Files, file)
	}
}

func (stat FileStats) ToString() string {
	return fmt.Sprintf("Type: %s; File Count: %d", stat.FileType.String(), len(stat.Files))
}

func CreateExtRegexp(ext string) *regexp.Regexp {
	reg, err := regexp.Compile(fmt.Sprintf(".*\\.%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	return reg
}

// *int means you *must* pass a *int (pointer to int), NOT just an int!
// func someFunc(x *int) {
// 	*x = 2 // Whatever variable caller passed in will now be 2
// 	y := 7
// 	x = &y // has no impact on the caller because we overwrote the pointer value!
// }

func main() {
	dir := "E:/Trevor/Videos/Radeon ReLive/Call of Duty Modern Warfare II  Warzone 2.0"

	image_stats := FileStats{image, CreateExtRegexp("png"), make([]fs.DirEntry, 1)}
	video_stats := FileStats{video, CreateExtRegexp("mp4"), make([]fs.DirEntry, 1)}
	audio_stats := FileStats{audio, CreateExtRegexp("m4a"), make([]fs.DirEntry, 1)}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		AppendIfMatch(&image_stats, file)
		AppendIfMatch(&video_stats, file)
		AppendIfMatch(&audio_stats, file)
	}

	fmt.Println(image_stats.ToString())
	fmt.Println(video_stats.ToString())
	fmt.Println(audio_stats.ToString())
}
