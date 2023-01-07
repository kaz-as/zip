package archive

import "io"

type File struct {
	In          io.Reader
	OutLocation []string
	OutName     string
}
