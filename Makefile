### Makefile --- 
## 
## Filename: Makefile
## Author: Mourad sabour
## Created: Mar mai  6 16:37:24 2014 (-0700)
## 

NAME=		gocss

BIN=		go

RM=		rm -rf

SRC=		src/main.go\
		src/parserCSS.go\
		src/utilities.go

all:
		$(BIN) build -o $(NAME) $(SRC)

clean:		
		$(RM) $(NAME) *~ *#

re:		
		make clean
		make all
