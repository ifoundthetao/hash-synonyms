package main
import (
    "crypto/md5"
    "fmt"
    "io/ioutil"
)

func checkForErrors(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fileContents, err := ioutil.ReadFile("./file.txt")
    checkForErrors(err)

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(fileContents))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
