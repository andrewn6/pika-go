package main

import ( 
  "fmt"
)
var BASE64_CHARS = []byte{
    'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
    'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
    'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
    'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
}

func Base64Encode(input string) string {
  var result []byte
  var buffer uint32 = 0
  var bitsLeft = 0

  for _, b := range []byte(input) {
    buffer = (buffer << 8) | uint32(b)
    bitsLeft += 8

    for bitsLeft >= 6 {
      bitsLeft -= 6
      index := int((buffer >> bitsLeft) & 0b111111)
      result = append(result, BASE64_CHARS[index])
    }
  }

  if bitsLeft > 0 {
    buffer <<= 6 - bitsLeft
    index := int(buffer & 0b111111)
    result = append(result, BASE64_CHARS[index])
  }

  for len(result) % 4 != 0 {
    result = append(result, '=')
  }

  return string(result)
}

func Base64Edecode(encoded string) ([]byte, error) {
  var decoded []byte
  var padding = 0
  var buffer uint32 = 0
  var bits = 0

  for _, c := range encoded {
    value := uint32(0)
    for i, char := range BASE64_CHARS {
      if char == byte(c) {
        value = uint32(i)
        break
      }
    }

    if value == 0 && c != 'A' {
      if padding > 0 {
        return nil, fmt.Errorf("invalid padding")
      }
      continue
    }

    buffer = (buffer << 6) | value
    bits += 6

    if bits >= 8 {
      bits -= 8
      decoded = append(decoded, byte(buffer>>bits))
      buffer &= (1 << bits) - 1
    }
  }

  if bits >= 6 || padding > 2 || (padding > 0 && bits > 0) {
    return nil, fmt.Errorf("invalid padding")
  }

  return decoded, nil
}
