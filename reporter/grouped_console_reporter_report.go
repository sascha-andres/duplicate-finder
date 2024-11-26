package reporter

import "path/filepath"

func (reporter *GroupedConsoleReporter) Report(files []string) {
	if len(files) <= 1 {
		return
	}

	reporter.fileList[len(reporter.fileList)] = files

	for outerIteration := range files {
		outerDir := filepath.Dir(files[outerIteration])
		for innerIteration := range files {
			if outerIteration == innerIteration {
				continue
			}
			innerDir := filepath.Dir(files[innerIteration])
			if _, ok := reporter.directoryFilesMap[outerDir]; ok {
				if innerValue, ok := reporter.directoryFilesMap[outerDir][innerDir]; ok {
					innerValue++
					reporter.directoryFilesMap[outerDir][innerDir] = innerValue
				} else {
					reporter.directoryFilesMap[outerDir][innerDir] = 1
				}
			} else {
				reporter.directoryFilesMap[outerDir] = make(map[string]int)
				reporter.directoryFilesMap[outerDir][innerDir] = 1
			}
		}
	}
	// get directory infos
}
