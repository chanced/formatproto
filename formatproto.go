// Package formatproto is a post-processor for protoc-gen-star that formats proto files with clang-format
package formatproto

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

type formatProto struct {
	currentFile string
}

// PostProcessor formats proto files with clang-format
func PostProcessor() pgs.PostProcessor {
	return &formatProto{}
}

func (p *formatProto) Match(a pgs.Artifact) bool {
	var n string
	switch a := a.(type) {
	case pgs.GeneratorFile:
		n = a.Name
	case pgs.GeneratorTemplateFile:
		n = a.Name
	case pgs.CustomFile:
		n = a.Name
	case pgs.CustomTemplateFile:
		n = a.Name
	default:
		return false
	}
	p.currentFile = n
	return strings.HasSuffix(n, ".proto")

}

func (p formatProto) Process(b []byte) ([]byte, error) {

	in := strings.NewReader(string(b))
	out := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	filename := path.Base(p.currentFile)

	cmd := exec.Command("clang-format", fmt.Sprintf("--assume-filename=%s", filename))
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = errbuf
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("%v\n %s", err, errbuf))
	}

	return out.Bytes(), nil
}

var _ pgs.PostProcessor = &formatProto{}
