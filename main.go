package main
import (
    "crypto/md5"
    "crypto/sha1"
    "flag"
    "fmt"
    "hash"
    "io/ioutil"
    "os"
)
const NO_ORIGINAL_INPUT_FILE_ERROR_NUMBER int = 1
const NO_MODDED_FILE_ERROR_NUMBER int = 2

func checkForErrors(e error) {
    if e != nil {
        fmt.Println("Error:", e)
        os.Exit(1)
    }
}

func getErrorMessages() []string {
    errorMessages := make([]string, 3)
    errorMessages[NO_ORIGINAL_INPUT_FILE_ERROR_NUMBER] = "No original input file provided"
    errorMessages[NO_MODDED_FILE_ERROR_NUMBER] = "No modified file provided"

    return errorMessages
}

func validateStartOptions() (*string, *string) {
    errorMessages := getErrorMessages()
    inputFileLocation := ""
    moddedFileLocation := ""

    flag.StringVar(&inputFileLocation, "include-file", "", "Original input file")
    flag.StringVar(&inputFileLocation, "i", inputFileLocation, "Original input file")

    flag.StringVar(&moddedFileLocation, "modified-file", "", "Modified file")
    flag.StringVar(&moddedFileLocation, "m", moddedFileLocation, "Modified file")

    flag.Parse()

    errorNumber := 0
    switch {
    case inputFileLocation == "":
        errorNumber = NO_ORIGINAL_INPUT_FILE_ERROR_NUMBER
    case moddedFileLocation == "":
        errorNumber = NO_MODDED_FILE_ERROR_NUMBER
    }
    hasError := errorNumber > 0
    if hasError {
        fmt.Println("Error:", errorMessages[errorNumber])
        os.Exit(errorNumber)
    }

    return &inputFileLocation, &moddedFileLocation
}

func getHasher(hashType string) hash.Hash {
    var newHash hash.Hash
    switch hashType{
    case "sha1":
        newHash = sha1.New()
    default:
        newHash = md5.New()
    }
    return newHash

}

func main() {
    var inputFileLocationPtr, moddedFileLocationPtr *string

    inputFileLocationPtr, moddedFileLocationPtr = validateStartOptions()
    fmt.Println(*moddedFileLocationPtr)

    fileContents, err := ioutil.ReadFile(*inputFileLocationPtr)
    checkForErrors(err)

    hasher := getHasher("sha1")

    hasher.Write([]byte(fileContents))
    hash := hasher.Sum(nil)
    hasher.Reset()
    fmt.Printf("%x\n", hash)
    hasher.Write([]byte("This is interesting"))
    hash = hasher.Sum(nil)
    hasher.Reset()
    fmt.Printf("%x\n", hash)
    hasher.Write([]byte(fileContents))
    hash = hasher.Sum(nil)
    hasher.Reset()
    fmt.Printf("%x\n", hash)
    hasher.Write([]byte(fileContents))
    hash = hasher.Sum(nil)
    hasher.Reset()
    fmt.Printf("%x\n", hash)


    //fmt.Printf("%x\n", md5Hash)
}
