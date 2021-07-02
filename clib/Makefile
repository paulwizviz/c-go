.Phony: all
all: libsum.a

libsum.a: sum.o
	ar -rsc libsum.a sum.o

sum.o: sum.c
	gcc -c sum.c -o sum.o

clean:
	rm libsum.a
	rm sum.o