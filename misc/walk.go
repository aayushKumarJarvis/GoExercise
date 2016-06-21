package misc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// WalkSourceDir checks the walking on the Directory
func WalkSourceDir(source, destination string) {
	filepath.Walk(source, func(path string, info os.FileInfo, walkErr error) error {
		fmt.Printf("path: %s\n", path)
		if info.Name() == "CURRENT" {
			dbLoc := filepath.Dir(path)
			dbBackupLoc := filepath.Join(destination, dbLoc)
			dbRelative, err := filepath.Rel(source, dbLoc)
			fmt.Printf("dbLoc : %s\n", dbLoc)
			fmt.Printf("dbRelative : %s\n", dbRelative)
			fmt.Printf("dbBackupLoc : %s\n", dbBackupLoc)
			if err != nil {
				log.Print(err)
				return err
			}
			return filepath.SkipDir
		}
		return nil
	})
}
