package reporter

import (
	"fmt"
	"io/ioutil"
)

func (reporter *GroupedConsoleReporter) Finish() error {
	fmt.Println("== potentially duplicated files ==")
	for i, files := range reporter.fileList {
		fmt.Println()
		fmt.Println(fmt.Sprintf("group %d:", i+1))
		for _, file := range files {
			fmt.Println(fmt.Sprintf("  %s", file))
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("== affected directories ==")
	for outerDir, commonFilesDirectory := range reporter.directoryFilesMap {
		fmt.Println()
		fmt.Println(fmt.Sprintf("Directory %s has files in common with:", outerDir))
		for innerDirectory, countOfFiles := range commonFilesDirectory {
			fileInfos, _ := ioutil.ReadDir(innerDirectory)
			fmt.Println(fmt.Sprintf("  - %s with %d file(s) with %d file(s) in the directory", innerDirectory, countOfFiles, len(fileInfos)))
		}
	}
	return nil
}
