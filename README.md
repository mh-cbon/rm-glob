# rm-glob

Rm globed files

## Install

Pick an msi package [here](https://github.com/mh-cbon/rm-glob/releases)!

__deb/rpm__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/rm-glob sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/rm-glob sh -xe
```

__go__

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon
cd $GOPATH/src/github.com/mh-cbon
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
