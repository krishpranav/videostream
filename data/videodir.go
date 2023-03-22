package data

type VideoDir struct {
	DirPath string
	Vidoes  []VideoInfo
}

type VideoDirectories []VideoDir

func (vd VideoDirectories) Len() int {
	return len(vd)
}

func (vd VideoDirectories) Swap(i, j int) {
	vd[i], vd[j] = vd[j], vd[i]
}

func NewVideoDir(path string, Videos []VideoInfo) VideoDir {
	return VideoDir{path, Videos}
}
