package main
import (
    "crypto/md5"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

func checkForErrors(e error) {
    if e != nil {
        fmt.Println("No File Given")
        os.Exit(1)
    }
}

func main() {
    var fileLocation string
    flag.StringVar(&fileLocation, "include-file", "Original input file", "")
    flag.StringVar(&fileLocation, "i", "Original input file", fileLocation)
    flag.Parse()

    fileContents, err := ioutil.ReadFile(fileLocation)
    checkForErrors(err)

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(fileContents))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
