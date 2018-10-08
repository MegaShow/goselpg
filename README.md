# Selpg

一个简单的CLI程序。

## Usage

```
Usage: selpg [-s startPage] [-e endPage] [-l linesPerPage | -f] filename

Options:
  -d, --dest string   destination
  -e, --end int       end page number (default max-page)
  -f, --formFeed      form feed per page
  -h, --help          help
  -l, --lines int     lines per page (default 72)
  -s, --start int     start page number (default 1)
```

Selpg没有强制选项，只包含可选选项和限制选项。

* `-sNumber`和`-eNumber`选项指定起始页数、终止页数，默认分别为首页、尾页。
* `-lNumber`选项指定每页行数，默认为72。
* `-f`选项指定页以换页符结束，而不是固定行数。
* `dDestination`选项指定将数据输出到打印机中。

其中，`-l`和`-f`为限制选项，两者只能同时指定一个。

下面为一些选项的简单用法。

```sh
$ selpg -s12 -e13 -l2

$ selpg -s12 -l2 input.txt

$ selpg -s12 -e12 -f input.txt

$ selpg -s12 -e12 -dlp1 input.txt
```

其中最后一条命令相当于：

```sh
$ selpg -s12 -e12 input.txt | lp -dlp1
```

## Design

除`main`外，本程序包含`command`和`print`两个包。

`command`用`pflag`实现，负责读取命令行选项参数并对参数进行检查。如果检查错误，将直接终止程序。该包不直接调用`print`的函数，而是返回选项参数数据，由`main`负责传参调用。

`print`负责接受选项参数数据，并根据参数决定如何处理数据。

## Check

程序对选项参数有部分检查。

如果指定非正数的起始页数、终止页数或行数，将返回错误。

```sh
$ selpg -s1 -e0 -l2
error: end page number cannot be non-positive
exit status 2
```

如果指定起始页数大于终止页数，将返回错误。

```sh
$ selpg -s2 -e1
error: start page number cannot be larger that end page number
exit status 2
```

如果同时指定每页行数和以换页符换页，将返回错误。

```sh
$ selpg -s1 -e2 -l2 -f
error: lines and formFeed cannot be set again
exit status 2
```

如果指定多个文件，将返回错误。

```sh
$ selpg -s1 -e2 -l2 input1.txt input2.txt
error: more than one file defined
exit status 2
```

## Test

下列命令可生成测试数据，并模拟三条命令的执行。

```sh
$ go test
```

测试代码生成文件`input.txt`，该文件有999行，每3行存在一个换页符。测试文件模拟命令如下：

```sh
$ go build selpg.go
$ ./selpg -s1 -e1 input.txt
$ ./selpg -s1 -e1 < input.txt
```

下列利用测试代码生成的文件`input.txt`，进行进一步测试。

将数据的第一页打印出来，每页5行。

```sh
$ selpg -s1 -e1 -l5 input.txt
lines 1
lines 2
lines 3
♀lines 4
lines 5
```

将重定向的数据第一页打印出来，每页5行。

```sh
$ selpg -s1 -e1 -l5 < input.txt
lines 1
lines 2
lines 3
♀lines 4
lines 5
```

将另一个命令的输出通过管道传输给Selpg，并将数据前两页打印出来，每页2行。

```sh
$ cat input.txt | selpg -s1 -e2 -l2
lines 1
lines 2
lines 3
♀lines 4
```

将Selpg的输出重定向到`output.txt`中。

```sh
$ selpg -s1 -e2 -l2 input.txt > output.txt
$ cat output.txt
lines 1
lines 2
lines 3
♀lines 4
```

将Selpg的输出丢弃。

```sh
$ selpg -s1 -e2 -l2 input.txt > /dev/null
```

将Selpg的输出通过管道传给另一个命令。

```sh
$ selpg -s1 -e2 -l2 input.txt | cat
lines 1
lines 2
lines 3
♀lines 4
```

将数据的前两页打印出来，每页以换页符标识结束。

```sh
$ selpg -s1 -e2 -f input.txt
lines 1
lines 2
lines 3
♀lines 4
lines 5
lines 6
♀
```

将Selpg的输出传输给打印机`lp1`。

```sh
$ selpg -s1 -e2 -l2 -dlp1 input.txt
```

## License

基于Apache License 2.0。

