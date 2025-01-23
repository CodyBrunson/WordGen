This is an updated remake of the [Babble](https://github.com/tjarratt/babble/) word generator. Not much has changed, it just detects the OS of the user and generates words appropriately.

The file that Unix has that is a long list of random words is embeded into the library and included along side the build.  This allows windows to access it.

### Dependencies
Requires the words file in the `/usr/share/dict` directory for non-windows users.

### Usage
```go
package your_app

import (
  "github.com/codybrunson/wordGen"
)

func main() {
  wordGen := NewWordGen()
  fmt.Println(wordGen.Generate())
})
```
By default *.Generate() returns 3 words separated by - (hyphen)

You can change anything about the generator by modifying the underlying wordGen options
wordGen.Seperator - This changes what the separator is between the words
wordGen.Count - How many words are generated any returned
wordGen.Words - Technically you can also change the word list to whatever you want as long as its the correct data type. []string

