# rm-glob

Rm globed files

## Install

Pick an msi package [here](https://github.com/mh-cbon/rm-glob/releases)!

__deb/ubuntu/rpm repositories__

```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/rm-glob sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/rm-glob sh -xe
```

__deb/ubuntu/rpm packages__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/rm-glob sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/rm-glob sh -xe
```

__chocolatey__

```sh
choco install rm-glob -y
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
