package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

func main() {
  dir, err := currentDirectory()
  if err != nil {
    fmt.Println("fatal error getting current working directory")
    return
  }
  var root_directory = flag.String("dir", dir, "the root directory to be scanned")
  flag.Parse()
  fmt.Println("Scanning directory: " + *root_directory)
  fileCounts := make(map[string]int)
  countFilesInDirectory(*root_directory, &fileCounts)
  totalFiles := 0
  for extension_type, fileCount := range fileCounts {
    fmt.Println(extension_type+":  ", fileCount)
    totalFiles += fileCount
  }
  fmt.Println("Total files scanned:  ", totalFiles)
}

func currentDirectory() (dir string, err error) {
  dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    return
  }
  fmt.Println("Current working directory is: " + dir)
  return
}

func countFilesInDirectory(dir string, fileCounts *map[string]int) {
  fileInfo, err := getDirInfo(dir)
  if err != nil {
    fmt.Println("error in countFilesInDirectory")
  }
  for _, file := range fileInfo {
    if file.IsDir() {
      dir_name := dir + "/" + file.Name()
      countFilesInDirectory(dir_name, fileCounts)
    } else {
      extension := getExtension(file.Name())
      (*fileCounts)[extension] += 1
    }
  }
}

func getExtension(fileName string) (extension string) {
  ar := strings.SplitAfterN(fileName, ".", 2)
  if len(ar) == 2 {
    extension = "." + ar[1]
  } else {
    extension = "No Extension"
  }
  return extension
}

func getDirInfo(dir string) (fileInfo []os.FileInfo, err error) {
  fileInfo, err = ioutil.ReadDir(dir)
  if err != nil {
    fmt.Println("fatal error reading directory")
  }
  return
}
