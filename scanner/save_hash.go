package scanner

// saveHash stores the hash into the map
func (scanner *Scanner) saveHash(path string, fileHash string) {
	scanner.duplicatesLock.Lock()
	defer scanner.duplicatesLock.Unlock()

	if val, ok := scanner.potentialDuplicates[fileHash]; ok {
		val = append(val, path)
		scanner.potentialDuplicates[fileHash] = val
	} else {
		val = make([]string, 1)
		val[0] = path
		scanner.potentialDuplicates[fileHash] = val
	}
}
