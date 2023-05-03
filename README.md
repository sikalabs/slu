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

Install latest version of **slu**  using [install-slu](https://github.com/sikalabs/install-slu) tool. See [install.sh](./install.sh)

```bash
curl -fsSL https://raw.githubusercontent.com/sikalabs/slu/master/install.sh | sudo sh
```

For upgrade of **slu** just run

```bash
install-slu install
```

### Windows (winget)

Using winget, official Windows package manager

```
winget install -e --id sikalabs.slu
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
