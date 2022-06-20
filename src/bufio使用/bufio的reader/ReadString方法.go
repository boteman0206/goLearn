package main

/**
底层调用ReadBytes
  func (b *Reader) ReadString(delim byte) (line string, err error) {
       bytes, err := b.ReadBytes(delim)
       return string(bytes), err
   }

*/
