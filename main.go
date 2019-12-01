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

func validateStartOptions() (*string, *string) {
    inputFileLocation := ""
    moddedFileLocation := ""

    flag.StringVar(&inputFileLocation, "include-file", "Original input file", "")
    flag.StringVar(&inputFileLocation, "i", "Original input file", inputFileLocation)

    flag.StringVar(&moddedFileLocation, "modified-file", "Modified file", "")
    flag.StringVar(&moddedFileLocation, "m", "Modded file", moddedFileLocation)

    flag.Parse()

    return &inputFileLocation, &moddedFileLocation
}


func main() {
    var inputFileLocationPtr, moddedFileLocationPtr *string

    inputFileLocationPtr, moddedFileLocationPtr = validateStartOptions()
    fmt.Println(*moddedFileLocationPtr)

    fileContents, err := ioutil.ReadFile(*inputFileLocationPtr)
    checkForErrors(err)

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(fileContents))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
