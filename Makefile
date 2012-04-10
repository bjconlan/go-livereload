# Looked into gyp, can't be bothered using it but if someone out there
# cares I like the idea of having this 'lil' app compile on windows which
# given the state of this Makefile requires some work...

BUILDDIR=out
CC=6g
LD=6l

# Order specific!
OBJS=${BUILDDIR}/exp/fsnotify/fsnotify.6 \
     ${BUILDDIR}/livereload.6

all: livereload

init:
	mkdir -p ${BUILDDIR}/exp/fsnotify

${BUILDDIR}/exp/fsnotify/fsnotify.6: init
	$(CC) -o ${BUILDDIR}/exp/fsnotify.6 src/exp/fsnotify/fsnotify.go src/exp/fsnotify/fsnotify_bsd.go

${BUILDDIR}/%.6 : src/%.go ${BUILDDIR}/exp/fsnotify/fsnotify.6
	$(CC) -o $@ -I ${BUILDDIR} $<

livereload: ${OBJS}
	$(LD) -o livereload -L ${BUILDDIR} ${BUILDDIR}/livereload.6

clean:
	rm -rf ${BUILDDIR} livereload
