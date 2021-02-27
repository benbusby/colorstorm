CC=gcc
CFLAGS=-O1 -I.

themes: themes.c
	$(CC) -o out themes.c $(CFLAGS)

clean:
	rm -f out
