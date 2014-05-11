# INC=-I/usr/include
LINK=-L/usr/lib -lglut -lGL -lGLU -lGLEW -lglfw -lm

all: main

run: all
	./main

clean:
	rm *.o

main: main.o modern.o
	gcc main.o modern.o -o main $(LINK)

main.o: main.c
	gcc $(INC) -c -o main.o main.c

modern.o: modern.c modern.h
	gcc $(INC) -c -o modern.o modern.c
