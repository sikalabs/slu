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

### Linux AMD64

```bash
# Check the current version on Github https://github.com/sikalabs/slu/releases
SLU_VERSION=v0.13.0 && \
SLU_ARCH=amd64 && \
curl -L https://github.com/sikalabs/slu/releases/download/${SLU_VERSION}/slu_${SLU_VERSION}_linux_${SLU_ARCH}.tar.gz \
  -o slu_${SLU_VERSION}_linux_${SLU_ARCH}.tar.gz && \
tar -xvzf slu_${SLU_VERSION}_linux_${SLU_ARCH}.tar.gz && \
rm slu_${SLU_VERSION}_linux_${SLU_ARCH}.tar.gz && \
mv slu /usr/local/bin/
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
