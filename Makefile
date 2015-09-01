NAME=peekaboo
VERSION=0.1
RELEASE:=$(shell date -u +%Y%m%d%H%M)
SRCDIR=src/github.com/mickep76/go-peekaboo
TMPDIR=.build
ARCH:=$(shell uname -p)
SRCDIR=src/github.com/mickep76/go-peekaboo/

all: build

clean:
	rm -f *.rpm
	rm -rf pkg bin ${TMPDIR}

test: clean
	gb test

build: test
	gb build all

update:
	gb vendor update --all

hwinfo:
	gb vendor update github.com/mickep76/hwinfo

pre-req:
	yum install -y rpm-build

rpm:	build
	mkdir -p ${TMPDIR}/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}
	cp -r bin ${SRCDIR}/templates ${SRCDIR}/public ${TMPDIR}/SOURCES
	sed -e "s/%NAME%/${NAME}/g" -e "s/%VERSION%/${VERSION}/g" -e "s/%RELEASE%/${RELEASE}/g" \
		${NAME}.spec >${TMPDIR}/SPECS/${NAME}.spec
	rpmbuild -vv -bb --target="${ARCH}" --clean --define "_topdir $$(pwd)/${TMPDIR}" ${TMPDIR}/SPECS/${NAME}.spec
	mv ${TMPDIR}/RPMS/${ARCH}/*.rpm .
