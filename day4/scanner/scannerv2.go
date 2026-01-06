package scanner

type ScannerV2 struct {
	bbox  BoundingBox
	input [][]string
}

func NewScannerV2(input [][]string) ScannerV2 {
	return ScannerV2{
		bbox: BoundingBox{
			lBox: 0,
			rBox: len(input[0]) - 1,
			tBox: 0,
			bBox: len(input) - 1,
		},
		input: input,
	}
}

func (s *ScannerV2) Scan() int {
	total := 0
	// If A is found && within bbox,
	// 	check NW to SE has M & S
	// 	check NE to SW has M & S
	// Can autotrim from the bbox since a cross needs 1 position from centre point to be valid
	for r := 1; r <= s.bbox.bBox-1; r++ {
		for c := 1; c < s.bbox.rBox; c++ {
			cursor := s.input[r][c]
			if cursor == "A" {
				// check side 1
				check1 := s.CheckBackSlash(r, c)
				// check side 2
				check2 := s.CheckForwardSlash(r, c)
				// if both true, ++
				if check1 && check2 {
					total++
				}
			}
		}
	}
	return total
}

func (s *ScannerV2) InBbox(r int, c int) bool {
	if r-1 < s.bbox.tBox {
		return false
	}
	if r+1 > s.bbox.bBox {
		return false
	}
	if c-1 < s.bbox.lBox {
		return false
	}
	if c+1 > s.bbox.rBox {
		return false
	}
	return true
}

func (s *ScannerV2) CheckBackSlash(r int, c int) bool {
	// M..
	// .A.
	// ..S
	if !s.InBbox(r, c) {
		return false
	}
	nw := s.input[r-1][c-1]
	se := s.input[r+1][c+1]
	if nw != "M" && nw != "S" {
		return false
	}
	if se != "M" && se != "S" {
		return false
	}

	if nw == "M" {
		return se == "S"
	} else {
		return se == "M"
	}
}

func (s *ScannerV2) CheckForwardSlash(r int, c int) bool {
	// ..S
	// .A.
	// M..
	if !s.InBbox(r, c) {
		return false
	}
	sw := s.input[r+1][c-1]
	ne := s.input[r-1][c+1]
	if sw != "M" && sw != "S" {
		return false
	}
	if ne != "M" && ne != "S" {
		return false
	}

	if sw == "M" {
		return ne == "S"
	} else {
		return ne == "M"
	}

}
