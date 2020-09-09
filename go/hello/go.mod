module hello

go 1.15

//replace github.com/randyzingle/sandbox/go/greetings => ../greetings

require github.com/randyzingle/sandbox/go/greetings v0.0.0-20200908204325-1356ee44d913 // indirect
