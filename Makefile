
all: clean doc compile

compile:
	./script/compile.sh
doc:
	./script/doc.sh
clean:
	./script/clean.sh