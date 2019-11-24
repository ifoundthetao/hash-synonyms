# Hash Synonyms
## Overview
Hash Synonyms is designed to find two different data representations that hash 
to the same value.  Similar to how two different words mean the same thing.

Hash Synonyms will take two input files:

    Original Input File -   This is the file which contains the original data
                            to generate the hash from

    Modified Input File -   This is a modified version of the input file to
                            to show where changes can be made to the file.
                            Changes are expressed by delimeters

# Use
`./hash-synonym -i original.file -m modified.file -d ^ -s % -l fixed -b"\x00\x0a"`

    -i, --include-file:     Original input file
    -m, --modified-file:    Modified input file
    -d, --delimeter-group:  Marker for the start/end of modifiable characters
    -s, --single-delimeter: Marker for a single character that can be modified
    -b, --bad-characters:   String with Hex values of characters not to use
    -f, --fixed:            Fixed length, output should be same length as the
                            input file.  The default is no restrictions.
    -o, --output-file       Location to write to, instead of stdout
