NAME:=$(shell basename `git rev-parse --show-toplevel`)
SRCDIR:=src/$(shell git config --get remote.origin.url | awk -F '[@:]' '{print $$2"/"$$3}' | sed 's/.git$$//')
TMPDIR=.build
VERSION:=$(shell awk -F '"' '/Version/ {print $$2}' ${SRCDIR}/version.go)
#RELEASE:=$(shell git rev-parse --verify --short HEAD)
RELEASE:=$(shell date -u +%Y%m%d%H%M)
ARCH:=$(shell uname -p)

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
		rpm.spec >${TMPDIR}/SPECS/${NAME}.spec
	rpmbuild -vv -bb --target="${ARCH}" --clean --define "_topdir $$(pwd)/${TMPDIR}" ${TMPDIR}/SPECS/${NAME}.spec
	mv ${TMPDIR}/RPMS/${ARCH}/*.rpm .
