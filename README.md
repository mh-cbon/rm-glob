# rm-glob

Rm globed files

# Install

You can grab a pre-built binary file in the [releases page](/mh-cbon/rm-glob/releases) 

```
mkdir -p $GOPATH/github.com/mh-cbon
cd $GOPATH/github.com/mh-cbon
git clone https://github.com/mh-cbon/rm-glob.git
cd rm-glob
glide install
go install
```

# Usage

```
NAME:
   rm-glob - Delete globed files

USAGE:
   rm-glob <options> <pattern>

VERSION:
   0.0.0

COMMANDS:
GLOBAL OPTIONS:
   --exclude value, -e value   Pattern of files to exclude
   --dry, -d                   Print files only, do not delete
   --recurse, -r               Recursively lookup directories
   --help, -h                  show help
   --version, -v               print the version

EXAMPLES:
  rm-glob -d "build"
  rm-glob -d "b**"
  rm-glob -d -r -e "some*" "bu*/**"
```
