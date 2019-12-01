package main
import (
    "crypto/md5"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)
const NO_ORIGINAL_INPUT_FILE_ERROR_NUMBER int = 1
const NO_MODDED_FILE_ERROR_NUMBER int = 2

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


func main() {
    var inputFileLocationPtr, moddedFileLocationPtr *string

    inputFileLocationPtr, moddedFileLocationPtr = validateStartOptions()
    fmt.Println(*moddedFileLocationPtr)

    fileContents, err := ioutil.ReadFile(*inputFileLocationPtr)

    md5Hasher := md5.New()

    md5Hasher.Write([]byte(fileContents))
    md5Hash := md5Hasher.Sum(nil)

    fmt.Printf("%x\n", md5Hash)
}
