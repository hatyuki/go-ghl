# ghl
ghl - GitHub Release Asset Locator

## Description
Get the asset location from GitHub releases

## Installation
Using `go get`:

```bash
go get github.com/hatyuki/go-ghl
```

## Usage
```bash
ghl [--os=OS] [--arch=ARCH] [--token=TOKEN] <Owner>/<Repository>[@<Tag>]
```

### Options
|Option |Description                                                                                     |
|-------|------------------------------------------------------------------------------------------------|
|--os   |Specifies the name of the target OS like 'windows', 'darwin' or 'linux'. (default: runtime.GOOS)|
|--arch |Specifies the name of the target architecture like '386' or 'amd64'. (default: runtime.GOARCH)  |
|--token|GitHub API Token. If not given, `ghl` reads it from `$GITHUB_TOKEN`.                            |

### Examples
```bash
ghl tcnksm/ghr
ghl tcnksm/ghr@v0.5.3                 # specifies the release version
ghl tcnksm/ghr --os linux --arch 368  # specifies the name of the target OS and architecture

wget -O ghr.zip $(ghl tcnksm/ghr)     # download latest release asset
```

### GitHub API Token
SEE: https://github.com/tcnksm/ghr#github-api-token
