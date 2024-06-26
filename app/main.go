package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	// filter pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	// flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	// order flags
	// hasOrderTime := flag.Bool("t", false, "sort by time, oldest to first")
	// hasOrderSize := flag.Bool("s", false, "sort by file size, smallest to bigger")
	// hasOrderReverse := flag.Bool("r", false, "reverse order sorting")

	flag.Parse()
	path := flag.Arg(0)

	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}
	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}

		isMatched, err := regexp.MatchString("(?i)"+*flagPattern, f.name)
		if err != nil {
			panic(err)
		}

		if !isMatched {
			continue
		}

		fs = append(fs, f)
	}
	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}

	printList(fs, *flagNumberRecords)

	// fmt.Println("Flag Pattern:", *flagPattern)
	// fmt.Println("Flag All:", *flagAll)
	// fmt.Println("Flag Number Records:", *flagNumberRecords)
	// fmt.Println("Flag Order Time:", *hasOrderTime)
	// fmt.Println("Flag Order Size:", *hasOrderSize)
	// fmt.Println("Flag Order Reverse:", *hasOrderReverse)
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]

		fmt.Printf("%s %s %s %10d %s %s %s %s\n",
			file.mode,
			file.userName,
			file.groupName,
			file.size,
			file.modificationTime.Format(time.DateTime),
			style.icon,
			file.name,
			style.symbol,
		)
	}
}

func getFile(dir fs.DirEntry, isHiding bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}
	f := file{
		name: dir.Name(),
		// fileType: 0,
		isDir:            dir.IsDir(),
		isHidden:         isHiding,
		userName:         "alex",
		groupName:        "primera-app",
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}
	setFile(&f)
	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.name, exe)
	}

	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpeg) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}
