package data

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoDirSort(t *testing.T) {
	dirsForTest := VideoDirectories{
		NewVideoDir("/home/krishpranav/Movies/", []VideoInfo{}),
		NewVideoDir("MEGAsync Downloads", []VideoInfo{}),
		NewVideoDir("/home/krishpranav/Downloads/", []VideoInfo{}),
	}
	sort.Sort(dirsForTest)
	for i := range dirsForTest {
		if i == len(dirsForTest)-1 {
			break
		}
		assert.Equal(t, dirsForTest[i].DirPath < dirsForTest[i+1].DirPath, true,
			fmt.Sprintf("Sort(VideoDirectories) return: %v", dirsForTest))
	}
}
