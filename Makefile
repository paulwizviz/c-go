.Phony: all
all: mylib/libsum.a
	cp clib/sum.h mylib/sum.h

mylib/libsum.a: clib/sum.o
	ar -rsc mylib/libsum.a clib/sum.o

clib/sum.o: clib/sum.c
	gcc -c clib/sum.c -o clib/sum.o

clean:
	rm mylib/libsum.a
	rm mylib/sum.h
	rm clib/sum.o