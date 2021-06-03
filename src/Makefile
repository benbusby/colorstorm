CC=gcc
CFLAGS=-O1 -I.
DEPS=themes.h
OBJ=themes.o main.o
OUT=earthbound-themes

%.o: %.c
	$(CC) -c -o $@ $< $(CFLAGS)

themes: $(OBJ)
	$(CC) -o $(OUT) $^ $(CFLAGS)

clean:
	rm -f *.o
