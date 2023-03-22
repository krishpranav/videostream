package filesys

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/krishpranav/videostream/data"
)

func generateVideoKey(videoPath string) string {
	hash := fnv.New32()
	hash.Write([]byte(videoPath))
	return fmt.Sprintf("%d%s", hash.Sum32(), filepath.Ext(videoPath))
}

func findVideosInPath(verbose bool, root string,
	videoDict data.VideoDict) []data.VideoDir {
	if verbose {
		log.Println("looking for videos in: " + root)
	}

	dirDict := make(map[string][]data.VideoInfo)

	filepath.Walk(root,
		func(pathStr string, info os.FileInfo, err error) error {
			if data.IsStreamableVideoFormat(path.Ext(pathStr)) {
				name := path.Base(pathStr)
				dir := filepath.Dir(pathStr)
				key := generateVideoKey(pathStr)
				newVideo := data.VideoInfo{
					FilePath: pathStr,
					FileName: name,
					Key:      key,
				}

				dirDict[dir] = append(dirDict[dir], newVideo)
				videoDict[key] = newVideo
			}
			return err
		})

	vids := []data.VideoDir{}
	for dirPath, videos := range dirDict {
		vids = append(vids, data.NewVideoDir(dirPath, videos))
	}
	return vids
}

func FindAllVideos(config data.Config) (data.VideoDirectories, data.VideoDict) {
	dirs := make(data.VideoDirectories, 0, 0)
	videoDict := make(map[string]data.VideoInfo)
	for _, pathStr := range config.VideoDirs {
		dirs = append(dirs, findVideosInPath(config.Verbose, pathStr, videoDict)...)
	}
	sort.Sort(dirs)
	return dirs, videoDict
}
