# Looked into gyp, can't be bothered using it but if someone out there
# cares I like the idea of having this 'lil' app compile on windows which
# given the state of this Makefile requires some work...

BUILDDIR=out
CC=6g
LD=6l

# Order specific!
OBJS=${BUILDDIR}/livereload/http_server.6 \
     ${BUILDDIR}/livereload.6

all: livereload

init:
	mkdir -p ${BUILDDIR}/livereload

${BUILDDIR}/%.6 : src/%.go init
	$(CC) -o $@ -I ${BUILDDIR} $<

livereload: ${OBJS}
	$(LD) -o livereload -L ${BUILDDIR} ${BUILDDIR}/livereload.6

clean:
	rm -rf ${BUILDDIR} livereload
