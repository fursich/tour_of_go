package main

import (
  "bytes"
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

var upperCase = []byte("abcdefghijklmnopqrstuvwxyz")
var lowerCase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (r *rot13Reader) Read(b []byte) (int, error) {

  reader := make([]byte, 1)
  _, err := r.r.Read(reader)

  if err != nil {
    return -1, io.EOF
  }

  ctr := 0
  var rot13 byte

  for i, c := range reader {
    pos1 := bytes.IndexByte(upperCase, c)
    pos2 := bytes.IndexByte(lowerCase, c)
    if pos1 != -1 {
      pos1 = (pos1 + 13) % 26
      rot13 = upperCase[pos1]
    } else if pos2 != -1 {
      pos2 = (pos2 + 13) % 26
      rot13 = lowerCase[pos2]
    } else {
      rot13 = c
    }
    ctr += 1
    b[i] = rot13
  }
  return ctr, nil
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
