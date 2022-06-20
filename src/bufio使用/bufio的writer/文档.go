package bufio的writer

/**
和 Reader 类型一样，bufio 包提供了两个实例化 bufio.Writer 对象的函数：NewWriter 和 NewWriterSize。其中，NewWriter 函数是调用 NewWriterSize 函数实现的：

// 实现了 io.ReaderFrom 接口
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// 实现了 io.Writer 接口
func (b *Writer) Write(p []byte) (nn int, err error)

// 实现了 io.ByteWriter 接口
func (b *Writer) WriteByte(c byte) error

// io 中没有该方法的接口，它用于写入单个 Unicode 码点，返回写入的字节数（码点占用的字节），内部实现会根据当前 rune 的范围调用 WriteByte 或 WriteString
func (b *Writer) WriteRune(r rune) (size int, err error)

// 写入字符串，如果返回写入的字节数比 len(s) 小，返回的error会解释原因
func (b *Writer) WriteString(s string) (int, error)


*/
