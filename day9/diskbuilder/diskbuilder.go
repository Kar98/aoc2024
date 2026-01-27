package diskbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

type Disk struct {
	startingBlock string
	disk          []FileBlock
}

type FileBlock struct {
	ID    int
	value int64
}

func NewDiskBuilder(startingBlock string) Disk {
	return Disk{
		startingBlock: startingBlock,
		disk:          make([]FileBlock, 0),
	}
}

func (d *Disk) Build() error {
	/// 12345
	// file, space, file, space, file
	// 0..111....22222

	file := true
	currentId := 0
	for char := range strings.SplitSeq(d.startingBlock, "") {
		if file {
			currentId++
		}
		if char == "0" {
			file = !file
			continue
		}
		num, err := strconv.ParseInt(char, 10, 64)
		if err != nil {
			return err
		}
		if file {
			for range num {
				block := FileBlock{ID: currentId, value: num}
				d.disk = append(d.disk, block)
			}
		} else {
			for range num {
				block := FileBlock{ID: -1, value: -1}
				d.disk = append(d.disk, block)
			}
		}
		file = !file
	}

	return nil
}

func (d *Disk) PrintDisk() {
	for _, block := range d.disk {
		fmt.Println(block)
	}
}
