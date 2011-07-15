## What Is This? ##
This project is a playground to familiarize myself with Google's Go programming language. The idea is to benchmark different Set implementation.

I wrote a blog post that explains all of this:
TODO

## How To Run It ##
There are 3 set implementations available. Modify the Makefile to include one of the implementation, remove the other lines (you can't have comments within a list like that) and then run `gotest -bench=. -memprofile=mem.out`. Then, pick a different implementation, run `make clean` then re-run the test.