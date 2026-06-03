#!/bin/sh

VERSION=v0.100.0 && \
OS=linux && \
ARCH=amd64 && \
BIN=slu && \
curl -fsSL https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz ${BIN} && \
rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
mv ${BIN} /usr/local/bin/slu && \
/usr/local/bin/slu install-bin slu
