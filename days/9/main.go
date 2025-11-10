package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
)

//go:embed input.txt
var rawInput string

func main() {
	lib.Day(
		input,
		part1,
		part2,
	)
}

type InputFile struct {
	Index int
	Size  int
}

func input() []InputFile {
	var files []InputFile

	file := true
	idx := 0
	for i := range strings.TrimSpace(rawInput) {
		d := lib.Atoi(rawInput[i : i+1])

		if file {
			files = append(files, InputFile{Index: idx, Size: d})
		}

		idx += d
		file = !file
	}

	return files
}

func part1(input []InputFile) uint64 {
	// Make mem (note: it's about 96kb), fill with -1
	lastFile := input[len(input)-1]
	mem := slices.Repeat([]int{-1}, lastFile.Index+lastFile.Size)

	// Write initial files
	for id, f := range input {
		for i := range f.Size {
			mem[f.Index+i] = id
		}
	}

	// Compress
	left := 0
	right := len(mem) - 1
	for {
		// Move cursors until left points to empty mem and right points to full mem
		for mem[left] >= 0 {
			left++
		}
		for mem[right] < 0 {
			right--
		}
		if left >= right {
			break
		}

		// Move
		mem[left] = mem[right]
		mem[right] = -1
	}

	// Checksum
	checksum := uint64(0)
	for idx, fileID := range mem {
		if fileID < 0 {
			break
		}
		checksum += uint64(idx) * uint64(fileID)
	}

	return checksum
}

func part2(input []InputFile) uint64 {
	// Make mem (note: it's about 96kb), fill with -1
	lastFile := input[len(input)-1]
	mem := slices.Repeat([]int{-1}, lastFile.Index+lastFile.Size)

	// Write initial files
	for id, f := range input {
		for i := range f.Size {
			mem[f.Index+i] = id
		}
	}

	// Compress
	// TODO: this is destructive, but we already ran part1 so it's fine
	slices.Reverse(input)
	for id, file := range input {
		id = len(input) - id - 1

		// Find gaps
		for i := range file.Index {
			if mem[i] >= 0 {
				continue
			}

			// We found a gap; how big is it?
			gapSize := 0
			for j := i; j < len(mem) && mem[j] < 0; j++ {
				gapSize++
			}

			// If it fits the file, move it
			if gapSize >= file.Size {
				for j := range file.Size {
					mem[i+j] = id
					mem[file.Index+j] = -1
				}
				break
			}

			// Otherwise, advance until after this gap (so we don't try to check it multiple times)
			i += gapSize - 1
		}
	}

	// Checksum
	checksum := uint64(0)
	for idx, fileID := range mem {
		if fileID < 0 {
			continue
		}
		checksum += uint64(idx) * uint64(fileID)
	}

	return checksum
}
