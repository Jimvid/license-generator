# License generator

## Overview

A CLI written in Go to generate licenses for your projects.
While GitHub offers an excellent GUI for adding licenses to projects,
It's just that... an GUI. It entails navigating to GitHub,
creating a file, inputting 'LICENSE', selecting a license, pushing it,
and subsequently pulling it locally. With this little tool, you can effortlessly
generate the license on your local machine and push it to GitHub.

## Installation

with Homebrew:

```bash
$ brew tap jimvid/jimvid && brew install license-generator
```

Binaries can also be installed manually, by downloading a zip file [here](https://github.com/Jimvid/license-generator/releases). These packages contain just a single executable file. You will have to set the executable bit on macOS and Linux.

## Usage

```bash
$ generate-licenses
```
