package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/skttomelo/file-stats/stats"
)

func CreateExtRegexp(ext string) *regexp.Regexp {
	reg, err := regexp.Compile(fmt.Sprintf(".*\\.%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	return reg
}

func main() {
	dir := "E:/Trevor/Videos/Radeon ReLive/Call of Duty Modern Warfare II  Warzone 2.0"

	image_stats := stats.DirStats{FileType: stats.IMAGE, Regexp: CreateExtRegexp("png"), DirEntries: make([]stats.DirEntryExt, 0)}
	video_stats := stats.DirStats{FileType: stats.VIDEO, Regexp: CreateExtRegexp("mp4"), DirEntries: make([]stats.DirEntryExt, 0)}
	audio_stats := stats.DirStats{FileType: stats.AUDIO, Regexp: CreateExtRegexp("m4a"), DirEntries: make([]stats.DirEntryExt, 0)}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		image_stats.AppendMatched(stats.CreateDirEntryExt(file, image_stats.FileType))
		video_stats.AppendMatched(stats.CreateDirEntryExt(file, video_stats.FileType))
		audio_stats.AppendMatched(stats.CreateDirEntryExt(file, audio_stats.FileType))
	}

	fmt.Println(image_stats.String())
	fmt.Println(video_stats.String())
	fmt.Println(audio_stats.String())
}
