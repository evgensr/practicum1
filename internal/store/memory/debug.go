package memory

import "github.com/davecgh/go-spew/spew"

func (box *Box) Debug() {
	spew.Dump(box)
}
