package stats

import (
	"fmt"
	"io/fs"
	"regexp"
	"time"
)

//go:generate stringer -type=FType
type FType uint8

const (
	IMAGE FType = iota
	VIDEO
	AUDIO
)

type DirEntryExt struct {
	DirEntry     fs.DirEntry // almost like we extending the functionality
	FileType     FType
	CreationTime time.Time // i'm going insane :killme:
}

type DirStats struct {
	FileType   FType
	Regexp     *regexp.Regexp
	DirEntries []DirEntryExt
}

func (ds *DirStats) AppendMatched(dee DirEntryExt) {
	if ds.Regexp.MatchString(dee.DirEntry.Name()) {
		ds.DirEntries = append(ds.DirEntries, dee)
	}
}

func (ds *DirStats) String() string {
	return fmt.Sprintf("Type: %s; File Count: %d", ds.FileType.String(), len(ds.DirEntries))
}
