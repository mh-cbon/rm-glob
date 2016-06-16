package main

import (
  "fmt"
  "os"
  "sort"
  "regexp"
  "strings"
  "io/ioutil"

	"github.com/mattn/go-zglob"
	"github.com/urfave/cli"
)

var VERSION = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "rm-glob"
	app.Version = VERSION
	app.Usage = "Delete globbed files"
	app.UsageText = "rm-glob <options> <pattern>"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "exclude, e",
      Value: "",
      Usage: "Pattern of files to exclude",
    },
    cli.BoolFlag{
      Name: "dry, d",
      Usage: "Print files only, do not delete",
    },
    cli.BoolFlag{
      Name: "recurse, r",
      Usage: "Recursively lookup directories",
    },
  }
  app.Action = func(c *cli.Context) error {

    pattern := string(c.Args().Get(0))
    exclude := c.String("exclude")
    dry     := c.Bool("dry")
    recurse := c.Bool("recurse")

    if len(pattern)==0 {
      return cli.NewExitError("pattern argument is required", 1)
    }

    matches, err := zglob.Glob(pattern)
    if err!=nil {
      return cli.NewExitError(err.Error(), 0) // soft failure
    }
    if recurse {
      matches = expandDirectories(matches)
    }

    if len(matches)==0 {
      return cli.NewExitError("pattern did not match any file", 0) // soft failure
    }

    okMatches := make([]string, 0)
    if len(exclude)>0 {
      excludeRe, err := getExlcudeRe(exclude, false)
      if err!=nil {
        return cli.NewExitError(err.Error(), 1)
      }

      for _, m := range matches {
        if excludeRe.MatchString(m)==false {
          okMatches = append(okMatches, m)
        }
      }
      if len(okMatches)==0 {
        return cli.NewExitError("pattern did not match any file after exclude pattern was applied", 0) // soft failure
      }
    } else {
      okMatches = append(okMatches, matches...)
    }

    sortByLength(&okMatches, "desc")
    if dry {
      for _, m := range okMatches {
        fmt.Println(m)
      }
    } else {
      for _, m := range okMatches {
        if _, err := os.Stat(m); !os.IsNotExist(err) {
          os.Remove(m)
        }
      }
    }

    return nil
  }

	app.Run(os.Args)
}

func getExlcudeRe(exclude string, sensitive bool) (*regexp.Regexp, error) {
	exclude = strings.Replace(exclude, "**", ".+", -1)
	exclude = strings.Replace(exclude, "*", "[/\\]+", -1)
	flags := ""
	if sensitive == false {
		flags = "(?i)"
	}
	return regexp.Compile(flags + exclude)
}

func expandDirectories(paths []string) []string {
  ret := paths
  for _, m := range paths {
    if stat, err := os.Stat(m); err==nil {
      if stat.IsDir() {
        ret = append(ret, expandDirectory(m)...)
      }
    }
  }
  return ret
}

func expandDirectory (dir string) []string {
  ret := make([]string, 0)
  files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ret
	}

	for _, file := range files {
    p := dir +"/" +file.Name()
    if stat, err := os.Stat(p); err==nil {
      if stat.IsDir() {
        ret = append(ret, p)
        ret = append(ret, expandDirectory(p)...)
      } else {
        ret = append(ret, p)
      }
    }
	}
  return ret
}

type AscStringList []string
func (s AscStringList) Len() int {
    return len(s)
}
func (s AscStringList) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s AscStringList) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}
type DescStringList []string
func (s DescStringList) Len() int {
    return len(s)
}
func (s DescStringList) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s DescStringList) Less(i, j int) bool {
    return len(s[i]) > len(s[j])
}
func sortByLength (s *[]string, dir string) {
  if dir=="asc" || dir=="ASC" {
    sort.Sort(AscStringList(*s))
  } else {
    sort.Sort(DescStringList(*s))
  }
}
