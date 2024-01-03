package stats

import (
	"io/fs"
	"log"
	"time"
)

func CreateDirEntryExt(entry fs.DirEntry, fType FType) DirEntryExt {
	return DirEntryExt{entry, fType, getCreationTime(entry)}
}

func getCreationTime(f fs.DirEntry) time.Time {
	info, err := f.Info()
	if err != nil {
		log.Fatal(err)
	}
	return info.ModTime()
}
