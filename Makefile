GC := go
OUTPUT := main.out
INPUT := main.go

build:
	cd src &&\
	$(GC) build -o $(OUTPUT) $(INPUT) &&\
	mv $(OUTPUT) ..
