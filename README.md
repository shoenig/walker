walker
======

A simple example of using filepath.Walk and WalkFunc. 

Demonstrates that the path argument is the *relative* path, and os.FileInfo.Name() is the *basename* of a file.
The path argument is only the same as the full path if the original root argument to filepath.Walk started with /.

```
$ walker .
path: .gitignore FileInfo.Name: .gitignore
path: README.md FileInfo.Name: README.md
path: main.go FileInfo.Name: main.go
path: sample FileInfo.Name: sample
path: sample/hello.txt FileInfo.Name: hello.txt
path: walker FileInfo.Name: walker
```

```
$ walker /tmp/foo
$ ./walker /tmp/foo
path: /tmp/foo FileInfo.Name: foo
path: /tmp/foo/abc FileInfo.Name: abc
path: /tmp/foo/abc/bar.txt FileInfo.Name: bar.txt
path: /tmp/foo/baz.txt FileInfo.Name: baz.txt
```
