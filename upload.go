package main

import (
  "fmt"
  "os"
  "code.google.com/p/ftp4go"
)

func main() {
  ftpClient := ftp4go.NewFTP(1) // 1 for debugging
  
  _, err := ftpClient.Connect("ftphost", ftp4go.DefaultFtpPort, "")
  if err != nil {
    fmt.Println("The connection failed.")
    os.Exit(1)
  }

  defer ftpClient.Quit()

  _, err = ftpClient.Login("username", "password", "")
  if err != nil {
    fmt.Println("Cannot login.")
    os.Exit(1)
  }

  // Change directory
  _, err = ftpClient.Cwd("/files")
  
  // List files in current directory
  var files []string
  files, err = ftpClient.Nlst()
  if err != nil {
    fmt.Println("Nlst failed.")
    os.Exit(1)
  }
  fmt.Println("Current folder contents:", files)

  // Upload a file
  err = ftpClient.UploadFile("test.txt", "/Users/jackie/test.txt", true, nil)
  if err != nil {
    fmt.Println("Failed to upload file.")  
  }
  
}
