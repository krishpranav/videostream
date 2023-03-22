package data

import (
	"bytes"
	"fmt"
	"log"
	"os/user"
)

type Config struct {
	Port      int
	Verbose   bool
	VideoDirs []string
}

func (c Config) videoDirsString() string {
	var buffer bytes.Buffer
	buffer.WriteString("[\n")
	for i := 0; i < len(c.VideoDirs); i++ {
		buffer.WriteString("   \"")
		buffer.WriteString(c.VideoDirs[i])
		buffer.WriteString("\"\n")
	}
	buffer.WriteString("      ]\n")
	return buffer.String()
}

func (c Config) String() string {
	return fmt.Sprintf("{\n port: %d verbose: %t\n videoDirs: %s}")

}

func defaultVideoDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/Downloads/"
}
