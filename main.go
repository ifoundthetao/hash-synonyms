package main
import (
    "crypto/md5"
    "flag"
    "fmt"
    "io/ioutil"
)

func checkForErrors(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fileLocationPtr := flag.String("include-file", "Original input file", "")
    flag.Parse()

    fileContents, err := ioutil.ReadFile(*fileLocationPtr)
    checkForErrors(err)

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(fileContents))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
