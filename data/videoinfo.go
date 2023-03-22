package data

type VideoInfo struct {
	FilePath string
	FileName string
	Key      string
}

type VideoDict map[string]VideoInfo

var videoFormats = []string{".avi", ".mkv", "mp4"}

func IsStreamableVideoFormat(videoExtension string) bool {
	for _, format := range videoFormats {
		if format == videoExtension {
			return true
		}
	}

	return false
}
