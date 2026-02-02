package diskbuilder

import (
	"errors"
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
	lookup map[int]FileBlock
}

type FileBlock struct {
	ID       int
	startPos int
	size     int
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
		Disk: Disk{startingBlock: startingBlock,
			disk: make([]FileBlock, 0)},
		lookup: make(map[int]FileBlock, 0),
	}
}

func (d *ComplexDisk) BuildDisk() error {
	file := true
	currentId := int(-1)
	cursor := 0
	for _, char := range strings.Split(d.startingBlock, "") {
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
			d.lookup[currentId] = FileBlock{ID: currentId, startPos: cursor, size: int(num)}
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
		cursor += int(num)
		file = !file
	}
	return nil
}

func (d *SimpleDisk) BuildDisk() error {

	file := true
	currentId := int(-1)
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

func (d *ComplexDisk) Sort() {
	// Get current file
	maxId := len(d.lookup) - 1
	// ID=0 will never be moved
	for i := maxId; i > 0; i-- {
		file := d.lookup[i]
		// Find an empty block of filesize
		newIdx, err := d.getEmptyBlock(file.startPos, file.size)
		if err != nil {
			// If not found, then go to the next file
			continue
		}
		// If found, then update d.disk
		for i := newIdx; i < newIdx+file.size; i++ {
			d.disk[i].ID = file.ID
		}
		// Remove old file
		for i := file.startPos; i < file.startPos+file.size; i++ {
			d.disk[i] = FileBlock{ID: -1}
		}
	}

}

func (d *Disk) CalculateChecksum() int64 {
	total := int64(0)
	for i, r := range d.disk {
		if d.disk[i].ID == -1 {
			continue
		}
		calc := int64(i) * int64(r.ID)
		total += calc
	}
	return total
}

func (d *ComplexDisk) getEmptyBlock(upToIdx int, size int) (int, error) {
	count := 0
	emptyIdx := 0
	for i := range upToIdx {
		block := d.disk[i]
		if block.ID != -1 {
			count = 0
			emptyIdx = i + 1
			continue
		}
		count++
		if count == size {
			return emptyIdx, nil
		}
	}
	return 0, errors.New("no block available")
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
