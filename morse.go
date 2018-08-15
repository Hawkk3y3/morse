package morse

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Encoder is an interface that specifies function to Encode
type Encoder interface {
	Encode(io.Reader) ([]byte, error)
}

// Decoder is an interface that specifies function to Encode
type Decoder interface {
	Decode(io.Writer, []byte) error
}

// Hacker is an interface that composes Encoder and Decoder
type Hacker interface {
	Encoder
	Decoder
}

type hacker struct {
}

func (h *hacker) Encode(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data := string(d)
	data = strings.ToUpper(data)
	var encodedValue string
	splitData := strings.Split(data, " ")
	numOfWords := len(splitData)
	for j, val := range splitData {
		vLen := len(val)
		for i, c := range val {
			char := string(c)
			encodedValue += alphaNumToMorse[char]
			if (i + 1) != vLen {
				encodedValue += " "
			}
		}
		if numOfWords != (j + 1) {
			encodedValue += " " + "/" + " "
		}
	}
	fmt.Println(encodedValue)
	return []byte(encodedValue), nil
}

func (h *hacker) Decode(w io.Writer, data []byte) error {
	return nil
}

// NewHacker is a factory function that generates a Morse Client
func NewHacker() Hacker {
	return &hacker{}
}
