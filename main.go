package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

func main() {
  dir, err := currentDirectory()
  if err != nil {
    fmt.Println("fatal error getting current working directory")
    return
  }
  var root_directory = flag.String("dir", dir, "the root directory to be scanned")
  flag.Parse()
  var fileNames []string
  fileNames = countFilesInDirectory(*root_directory, &fileNames)
  for _, file := range fileNames {
    fmt.Println(file)
  }
}

func currentDirectory() (dir string, err error) {
  dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    return
  }
  fmt.Println("Current working directory is: " + dir)
  return
}

func countFilesInDirectory(dir string, files *[]string) []string {
  fileInfo, err := getDirInfo(dir)
  if err != nil {
    fmt.Println("error in countFilesInDirectory")
    return *files
  }
  for _, file := range fileInfo {
    if file.IsDir() {
      dir_name := dir + "/" + file.Name()
      countFilesInDirectory(dir_name, files)
    } else {
      *files = append(*files, dir+"/"+file.Name())
    }
  }
  return *files
}

func getDirInfo(dir string) (fileInfo []os.FileInfo, err error) {
  fileInfo, err = ioutil.ReadDir(dir)
  if err != nil {
    fmt.Println("fatal error reading directory")
  }
  return
}
