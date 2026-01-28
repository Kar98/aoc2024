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

type SimpleDisk struct {
	Disk
}

type ComplexDisk struct {
	Disk
}

type FileBlock struct {
	ID       int64
	startPos int
	length   int
}

func NewSimpleDisk(startingBlock string) SimpleDisk {
	return SimpleDisk{
		Disk{
			startingBlock: startingBlock,
			disk:          make([]FileBlock, 0),
		},
	}
}

func NewComplexDisk(startingBlock string) ComplexDisk {
	return ComplexDisk{
		Disk{
			startingBlock: startingBlock,
			disk:          make([]FileBlock, 0),
		},
	}
}

func (d *ComplexDisk) BuildDisk() error {
	return nil
}

func (d *SimpleDisk) BuildDisk() error {

	file := true
	currentId := int64(-1)
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
				block := FileBlock{ID: currentId}
				d.disk = append(d.disk, block)
			}
		} else {
			for range num {
				block := FileBlock{ID: -1}
				d.disk = append(d.disk, block)
			}
		}
		file = !file
	}

	return nil
}

func (d *Disk) PrintDisk() string {
	var output strings.Builder
	for _, block := range d.disk {
		if block.ID == -1 {
			output.WriteString(".")
		} else {
			output.WriteString(fmt.Sprint(block.ID))
		}

	}
	return output.String()
}

func (d *SimpleDisk) Sort() {
	totalIterations := d.countEmptySpaces()
	totalLen := len(d.disk) - 1

	emptyIndex := 0
	// Iterate thorugh the loop and grab the last value and put it into an empty space
	for i := totalLen; i >= totalLen+1-totalIterations; i-- {
		if d.disk[i].ID == -1 {
			continue
		}
		emptyIndex = d.getEmptyBlock(emptyIndex)
		// Write to empty block
		d.disk[emptyIndex].ID = d.disk[i].ID
		// Update source block to be empty
		d.disk[i].ID = -1
	}
}

func (d *Disk) CalculateChecksum() int64 {
	total := int64(0)
	for i, r := range d.disk {
		if d.disk[i].ID == -1 {
			continue
		}
		calc := int64(i) * r.ID
		total += calc
	}
	return total
}

func (d *SimpleDisk) getEmptyBlock(fromIdx int) int {
	for i := fromIdx; i < len(d.disk)-1; i++ {
		if d.disk[i].ID == -1 {
			return i
		}
	}
	return -1 // should cause a panic
}

func (d *SimpleDisk) countEmptySpaces() int {
	total := 0
	for _, block := range d.disk {
		if block.ID == -1 {
			total++
		}
	}
	return total
}
