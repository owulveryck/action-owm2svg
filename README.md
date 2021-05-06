# Create a SVG representation of an onlinewardleymap


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