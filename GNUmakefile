VSN = 0.8
BIN = slf2tcx
REL = $(BIN)-$(GOHOSTOS)-$(GOHOSTARCH)-$(VSN)
SRCS = gpx.go main.go slf.go tcx.go

GOHOSTARCH = $(shell go env GOHOSTARCH)
GOHOSTOS = $(shell go env GOHOSTOS)

.PHONY: all clean distclean rel

all: $(BIN)

clean:
	rm -f $(BIN)

distclean: clean
	rm -f $(REL)

rel: $(REL)

$(BIN): $(SRCS)
	go build -o $@ $^

$(REL): $(BIN)
	cp $< $@
	strip $@
