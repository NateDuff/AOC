package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type File struct {
	ID   int
	Size int
}

func main() {
	input, _ := os.ReadFile("../input/day9.txt")
	diskmap := strings.TrimSpace(string(input)) + "0"

	files1, files2 := []File{}, []File{}
	// Parse diskmap into files1 and files2
	for id := 0; id*2 < len(diskmap); id++ {
		size, free := int(diskmap[id*2]-'0'), int(diskmap[id*2+1]-'0')
		// Add files to files1
		files1 = append(files1, slices.Repeat([]File{{id, 1}}, size)...)
		// Add free space to files1
		files1 = append(files1, slices.Repeat([]File{{-1, 1}}, free)...)

		// Add files to files2
		files2 = append(files2, File{id, size}, File{-1, free})
	}
	fmt.Println(getChecksumForFile(files1))
	fmt.Println(getChecksumForFile(files2))
}

// getChecksumForFile calculates the checksum for a given list of files.
func getChecksumForFile(files []File) (checksum int) {
	// Sort files by size
	for file := len(files) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			// If the file is not empty and the free space is empty and the free space is bigger than the file
			if files[file].ID != -1 && files[free].ID == -1 && files[free].Size >= files[file].Size {
				// Insert the file into the free space
				files = slices.Insert(files, free, files[file])
				files[file+1].ID, files[free+1].ID = -1, -1
				files[free+1].Size = files[free+1].Size - files[file+1].Size
			}
		}
	}

	// Calculate checksum
	i := 0
	for _, f := range files {
		for range f.Size {
			if f.ID != -1 {
				checksum += i * f.ID
			}
			i++
		}
	}
	return checksum
}
