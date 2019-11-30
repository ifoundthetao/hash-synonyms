package main
import (
    "crypto/md5"
    "fmt"
)

func main() {
    stringToHash := "foo bar this is hashed"

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(stringToHash))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
