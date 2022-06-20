package main

/*
ReadLine 是一个底层的原始行读取命令。许多调用者或许会使用 ReadBytes('\n') 或者 ReadString('\n') 来代替这个方法。


ReadLine 尝试返回单独的行，不包括行尾的换行符。如果一行大于缓存，isPrefix 会被设置为 true，同时返回该行的开始部分（等于缓存大小的部分）。
该行剩余的部分就会在下次调用的时候返回。当下次调用返回该行剩余部分时，isPrefix 将会是 false 。
跟 ReadSlice 一样，返回的 line 只是 buffer 的引用，在下次执行IO操作时，line 会无效。可以将 ReadSlice 中的例子该为 ReadLine 试试。


注意，返回值中，要么 line 不是 nil，要么 err 非 nil，两者不会同时非 nil。

*/
