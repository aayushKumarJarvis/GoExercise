package misc

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tecbot/gorocksdb"
)

// Current is used to identify the backup directory
const Current = "CURRENT"

// ReadTarballs file reads the .tar.gz files and returns its contents
func ReadTarballs() {
	flag.Parse() // get the arguments from command line

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		fmt.Println("Usage : practice $SOURCE_FILE")
		os.Exit(1)
	}

	file, err := os.Open(sourceFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourceFile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		case tar.TypeReg:
			// handle normal file
			fmt.Println("Extracting :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			writer.Close()
		default:
			fmt.Printf("Unable to extract type : %c in file %s", header.Typeflag, filename)
		}
	}
}

// CreateBackup function reads mpid_index_db and store_db directory from extracted tarball and triggers a backup for the same
func CreateBackup(source string) error {

	fmt.Printf("Trigerring Backup for %s\n", source)
	opts := gorocksdb.NewDefaultOptions()
	rocksdb, err := gorocksdb.OpenDb(opts, source)
	db, err := gorocksdb.OpenBackupEngine(opts, source)

	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	db.CreateNewBackup(rocksdb)
	return nil
}
