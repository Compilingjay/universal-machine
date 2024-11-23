# Universal Machine in Go
This is an implementation of the Universal Machine, based on a challenge from ICFP 2006. You can learn more about it here: [Link](http://boundvariable.org/task.shtml)

To run, build the project and run like so:
```
go build
./universal-machine <filepath>

```

If you would like to run the program embedded in the codex, you can use the included program.um under publications/, or build it yourself using the provided shell or batch files, and then filtering the beginning of the output to get your program by using a good text editor (I recommend nvim) like so:

```
go build
./<shell/batch file>
mv out.txt program.um
nvim program.um
./universal-machine program.um
```

The remaining programs and relevant files are included under publications/.