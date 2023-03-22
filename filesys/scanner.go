package filesys

import (
	"fmt"
	"hash/fnv"
	"path/filepath"
	"sort"

	"github.com/krishpranav/videostream/data"
)

func generateVideoKey(videoPath string) string {
	hash := fnv.New32()
	hash.Write([]byte(videoPath))

	return fmt.Sprintf("%d%s", filepath.Ext(videoPath))
}

func findVideoInPath(verbose bool, root string) {

}

func FindAllVideos(config data.Config) (data.VideoDirectories, data.VideoDict) {
	dirs := make(data.VideoDirectories, 0, 0)
	videoDict := make(map[string]data.VideoInfo)
	for _, pathStr := range config.VideoDirs {
		dirs = append(dirs, findVideoInPath(config.Verbose, pathStr, config.VideoDirs[])...)
	}

	sort.Sort(dirs)
	return dirs
}
