package main

/*
$ cp --help
Usage: cp [OPTION]... [-T] SOURCE DEST
  or:  cp [OPTION]... SOURCE... DIRECTORY
  or:  cp [OPTION]... -t DIRECTORY SOURCE...
Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.

Mandatory arguments to long options are mandatory for short options too.
  -a, --archive                same as -dR --preserve=all
      --attributes-only        don't copy the file data, just the attributes
      --backup[=CONTROL]       make a backup of each existing destination file
  -b                           like --backup but does not accept an argument
      --copy-contents          copy contents of special files when recursive
  -d                           same as --no-dereference --preserve=links
  -f, --force                  if an existing destination file cannot be
                                 opened, remove it and try again (this option
                                 is ignored when the -n option is also used)
  -i, --interactive            prompt before overwrite (overrides a previous -n
                                  option)
  -H                           follow command-line symbolic links in SOURCE
  -l, --link                   hard link files instead of copying
  -L, --dereference            always follow symbolic links in SOURCE
  -n, --no-clobber             do not overwrite an existing file (overrides
                                 a previous -i option)
  -P, --no-dereference         never follow symbolic links in SOURCE
  -p                           same as --preserve=mode,ownership,timestamps
      --preserve[=ATTR_LIST]   preserve the specified attributes (default:
                                 mode,ownership,timestamps), if possible
                                 additional attributes: context, links, xattr,
                                 all
      --no-preserve=ATTR_LIST  don't preserve the specified attributes
      --parents                use full source file name under DIRECTORY
  -R, -r, --recursive          copy directories recursively
      --reflink[=WHEN]         control clone/CoW copies. See below
      --remove-destination     remove each existing destination file before
                                 attempting to open it (contrast with --force)
      --sparse=WHEN            control creation of sparse files. See below
      --strip-trailing-slashes  remove any trailing slashes from each SOURCE
                                 argument
  -s, --symbolic-link          make symbolic links instead of copying
  -S, --suffix=SUFFIX          override the usual backup suffix
  -t, --target-directory=DIRECTORY  copy all SOURCE arguments into DIRECTORY
  -T, --no-target-directory    treat DEST as a normal file
  -u, --update                 copy only when the SOURCE file is newer
                                 than the destination file or when the
                                 destination file is missing
  -v, --verbose                explain what is being done
  -x, --one-file-system        stay on this file system
  -Z                           set SELinux security context of destination
                                 file to default type
      --context[=CTX]          like -Z, or if CTX is specified then set the
                                 SELinux or SMACK security context to CTX
      --help     display this help and exit
      --version  output version information and exit

By default, sparse SOURCE files are detected by a crude heuristic and the
corresponding DEST file is made sparse as well.  That is the behavior
selected by --sparse=auto.  Specify --sparse=always to create a sparse DEST
file whenever the SOURCE file contains a long enough sequence of zero bytes.
Use --sparse=never to inhibit creation of sparse files.

When --reflink[=always] is specified, perform a lightweight copy, where the
data blocks are copied only when modified.  If this is not possible the copy
fails, or if --reflink=auto is specified, fall back to a standard copy.
Use --reflink=never to ensure a standard copy is performed.

The backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.
The version control method may be selected via the --backup option or through
the VERSION_CONTROL environment variable.  Here are the values:

  none, off       never make backups (even if --backup is given)
  numbered, t     make numbered backups
  existing, nil   numbered if numbered backups exist, simple otherwise
  simple, never   always make simple backups

As a special case, cp makes a backup of SOURCE when the force and backup
options are given and SOURCE and DEST are the same name for an existing,
regular file.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation at: <https://www.gnu.org/software/coreutils/cp>
or available locally via: info '(coreutils) cp invocation'

*/

// import (
// 	"fmt"
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
// 	app := &cli.App{
// 		Name:  "cp",
// 		Usage: "cp summary info",
// 		Description: `
// cp detail info
// cp detail info2
// cp detail info3
// `,
// 		UseShortOptionHandling: true,
// 		Version:                "v0.0.1",
// 		ReleaseDate:            "2021-02-04",
// 		Copyright:              "copyright @some-INC",
// 		Authors: []*cli.Author{
// 			&cli.Author{"author1", "author1@some-inc.com"},
// 			&cli.Author{"author2", "author2@other-inc.com"},
// 		},
// 		Flags: []cli.Flag{
// 			&cli.BoolFlag{
// 				Name:    "verbose",
// 				Value:   true,
// 				Aliases: []string{"V"},
// 				Usage:   "verbose",
// 			},
// 			&cli.BoolFlag{
// 				Name:    "force",
// 				Value:   true,
// 				Aliases: []string{"f"},
// 				Usage:   "force",
// 			},
// 			&cli.BoolFlag{
// 				Name:     "recursive",
// 				Aliases:  []string{"R", "r"},
// 				Required: true,
// 				Usage:    "recursive",
// 			},
// 			&cli.StringFlag{
// 				Name:      "link",
// 				Aliases:   []string{"l"},
// 				LogicName: "url",
// 				Usage:     "link url",
// 				Value:     "default",
// 				EnvVars:   []string{"ENV_CP_LINK"},
// 			},
// 			&cli.StringSliceFlag{
// 				Name:      "0",
// 				LogicName: "SOURCE",
// 				Usage:     "source path",
// 				Value:     cli.NewStringSlice("default1", "default2"),
// 				EnvVars:   []string{"ENV_CP_0"},
// 			},
// 			&cli.Float64SliceFlag{
// 				Name:        "fs",
// 				LogicName:   "",
// 				Usage:       "Float64SliceFlag",
// 				Value:       cli.NewFloat64Slice(1.1, 2.2),
// 				DefaultText: "[1.1,2.2]",
// 			},
// 			&cli.Int64SliceFlag{
// 				Name:      "i64s",
// 				LogicName: "",
// 				Usage:     "Int64SliceFlag",
// 				Value:     cli.NewInt64Slice(1, 2),
// 			},
// 			&cli.IntSliceFlag{
// 				Name:      "is",
// 				LogicName: "",
// 				Usage:     "IntSliceFlag",
// 				Value:     cli.NewIntSlice(1, 2),
// 			},
// 			&cli.StringFlag{
// 				Name:      "1",
// 				LogicName: "DEST",
// 				Usage:     "destnation path",
// 			},
// 		},
// 		Action: func(ctx *cli.Context) error {
// 			fmt.Println(ctx.LocalFlagNames())
// 			fmt.Println(ctx.FlagNames())
// 			for i, v := range ctx.FlagNames() {
// 				fmt.Printf("%d %s %#v\n", i, v, ctx.Value(v))
// 			}
// 			fmt.Printf("%s %+v\n", "-0", ctx.Value("0"))
// 			err := ctx.Err()
// 			fmt.Println(err)
// 			if err == nil {
// 				cli.ShowAppHelp(ctx, nil)
// 			}
// 			return nil
// 		},
// 	}
// 	line := `cp -rf -V=false -v=0 --link="link url" -0=set13,set14 -fs=13.13 -i64s=13 -is=13 -1="/a"`
// 	args := os.Args
// 	if len(args) <= 1 {
// 		args = cli.SplitCommandLine(line)
// 	}
// 	fmt.Printf("args: %+v\n\n", args)
// 	os.Setenv("ENV_CP_LINK", "link_env")
// 	fmt.Println("ENV_CP_LINK", os.Getenv("ENV_CP_LINK"))
// 	os.Setenv("ENV_CP_0", "0_env")
// 	fmt.Println("ENV_CP_0", os.Getenv("ENV_CP_0"))
// 	if err := app.Run(args); err != nil {
// 		err = err
// 		//fmt.Println(err)
// 	}
// }

//package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	os.Args = []string{"multi_values",
		"--stringSclice", "parsed1,parsed2", "--stringSclice", "parsed3,parsed4",
		"--float64Sclice", "13.3,14.4", "--float64Sclice", "15.5,16.6",
		"--int64Sclice", "13,14", "--int64Sclice", "15,16",
		"--intSclice", "13,14", "--intSclice", "15,16",
	}
	os.Setenv("ENV_CLI_FLOAT64_SLICE", "23.3,24.4")
	app := cli.NewApp()
	app.Name = "multi_values"
	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{Name: "stringSclice"},
		&cli.Float64SliceFlag{Name: "float64Sclice", EnvVars: []string{"ENV_CLI_FLOAT64_SLICE"}},
		&cli.Int64SliceFlag{Name: "int64Sclice"},
		&cli.IntSliceFlag{Name: "intSclice"},
	}
	app.Action = func(ctx *cli.Context) error {
		for i, v := range ctx.FlagNames() {
			fmt.Printf("%d-%s %#v\n", i, v, ctx.Value(v))
		}
		return ctx.Err()
	}

	_ = app.Run(os.Args)
}
