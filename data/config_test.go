package data

import (
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
)

func TestParseJsonConfigNullConfig(t *testing.T) {
	_, err := ParseJsonConfig("")
	assert.NotNil(t, err)
}

func TestParseJsonConfig(t *testing.T) {
	defer filet.CleanUp(t)
	file := filet.TmpFile(t, "", `{
		"port": 5678,
		"verbose": true,
		"videoDirs": [ 
			"/home/krishpranav/Movies/",
			"/home/krishpranav/Downloads/" 
			"MEGAsync Downloads"
		]
	}
`)
	configForTest := Config{
		Port:    5678,
		Verbose: true,
		VideoDirs: []string{
			"/home/krishpranav/Downloads/",
			"/home/krishpranav/Movies/",
			"MEGAsync Downloads"},
	}
	c, err := ParseJsonConfig(file.Name())
	assert.Nil(t, err)
	assert.Equal(t, *c, configForTest, "ParseJsonConfig() return wrong config")
}

func TestParseJsonConfigEmpty(t *testing.T) {
	defer filet.CleanUp(t)
	file := filet.TmpFile(t, "", "")
	c, err := ParseJsonConfig(file.Name())
	assert.Nil(t, err)
	assert.Equal(t, c, DefaultConfig(), "ParseJsonConfig(Empty) should return default config")
}
