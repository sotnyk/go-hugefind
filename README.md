# go-hugefind
This is my first project on the GO language

It is a small command-line utility which counts how many times some substring is encountered in the text file. It is simple task but I can't find
such utility for huge files. For ex., we need to count how many products send us some shop. His feed is a XML with size more than 5 GB. We should to count entries of "**<product **" substring to answer on our question.

# Usage
```
go-findhuge substring filename
```

Substring can contain symbols which is reserved in your shell. We should escape it. How it can be made with Windows command-line, you can read [on this page](http://ss64.com/nt/syntax-esc.html). So, to count the "**<product **" substring we should use following command:

```
go-findhuge "^<product " feed_shop1.xml
```