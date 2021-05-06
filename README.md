# Create a SVG representation of an onlinewardleymap

This action creates a SVG representation of a [`owm`](https://onlinewardleymaps.com/) file.
The output has the same name of the input with a `svg` suffix.

## Inputs

### `owmfile`

**Required** The name of the input file in owm format.

## Outputs

### svgfile

The SVG filename

## Example usage

```yml
name: owm2svg example

on:
  push:
    branches:
    - main
  pull_request:
    branches:
    - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v1
    - uses: owulveryck/action-owm2svg
      with:
        owmfile: 'sample/test.owm'
    - name: Push to doc branch
      uses: Automattic/action-commit-to-branch@master
      with:
        branch: 'doc'
        commit_message: 'SVG Generation'
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }} 
```