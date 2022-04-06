<p align="center">
  <h1 align="center">slu: SikaLabs Utils</h1>
  <p align="center">
    <a href="https://opensource.sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/OPENSOURCE BY-SIKALABS-131480?style=for-the-badge"></a>
    <a href="https://sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-sikalabs.com-gray?style=for-the-badge"></a>
    <a href="mailto:opensource@sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-opensource@sikalabs.com-gray?style=for-the-badge"></a>
  </p>
</p>

## Install

### Mac

```
brew install sikalabs/tap/slu
```

### Linux AMD64 (using install-slu, recommended)

Install latest version of **slu**  using [install-slu](https://github.com/sikalabs/install-slu) tool.

```bash
sudo su -
VERSION=v0.1.0 && \
OS=linux && \
ARCH=amd64 && \
BIN=install-slu && \
curl -L https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xvzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
mv ${BIN} /usr/local/bin/ && \
install-slu install
```

You can do the same with this one-liners

```bash
sudo su - && VERSION=v0.1.0 && OS=linux && ARCH=amd64 && BIN=install-slu && curl -L https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && tar -xvzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && mv ${BIN} /usr/local/bin/ && install-slu install
```

```bash
curl -fsSL https://ins.oxs.cz/slu-linux-amd64.sh | sudo sh
```

For upgrade of **slu** just run

```bash
install-slu install
```

### Linux AMD64 (directly)

```bash
sudo su -
# Check the current version on Github https://github.com/sikalabs/slu/releases
SLU_VERSION=v0.38.0 && \
VERSION=$SLU_VERSION && \
OS=linux && \
ARCH=amd64 && \
BIN=slu && \
curl -L https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xvzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
mv ${BIN} /usr/local/bin/
```

### Windows (scoop)

Install unsing [scoop](https://scoop.sh/)

```
scoop install https://raw.githubusercontent.com/sikalabs/scoop-bucket/master/slu.json
```

### Autocomplete

See: `slu completion`

#### Bash

```
source <(slu completion bash)
```

## CLI Usage

### Generated CLI Docs on Github

See: <https://github.com/sikalabs/slu-cli-docs/blob/master/slu.md#slu>

## Generate CLI Docs

Generate Markdown CLI docs to `./cobra-docs`

```
slu generate-docs
```
