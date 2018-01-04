package main

import (
	"github.com/mobileka/glean/pkg/debug"
	"github.com/mobileka/glean/pkg/filesystem"
	"github.com/mobileka/glean/pkg/glean"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cfg = kingpin.Flag(
		"config",
		"a path to the configuration file. Should include the filename. Default: ./glean.yaml").
		Short('c').
		String()

	jpegDir = kingpin.Flag(
		"jd",
		"a directory where JPEG files are located. Default:the current directory").String()

	rawDir = kingpin.Flag(
		"rd",
		"a directory where RAW files are located. Default:the current directory").String()

	jpegExt = kingpin.Flag(
		"je",
		"a list of file extensions representing JPEG files. Default: [jpeg, jpg]").Strings()

	rawExt = kingpin.Flag(
		"re",
		"a list of file extensions representing RAW files. Default: [nef]").Strings()

	dbg = kingpin.Flag(
		"debug",
		"a list of file extensions representing RAW files. Default: [nef]").
		Default("false").Short('d').Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	serf := equipTheSerf()

	argv := glean.Arguments{JpegDir: *jpegDir, RawDir: *rawDir, JpegExt: *jpegExt, RawExt: *rawExt}

	debugger := debug.NewDebugger(*dbg)
	fs := filesystem.NewFilesystem()

	curDir, err := fs.GetCurrentDirectory()
	if err != nil {
		panic(err)
	}
	debugger.Logf("Current Directory: %s\n", curDir)

	config := glean.NewConfig(*cfg, curDir)

	serf.ExpressHowAmusedYouAreToStart()

	cmd := glean.NewGlean(debugger, config, fs, argv)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	serf.ExpressHowHappyYouAreWithResults()
}

func equipTheSerf() *glean.Serf {

	start := []string{
		"Why am I doing all de hard work?",
		"Happily killing your files with 🔥",
		"I'm not crazy about gleaning your shit, ya know?",
		"You're so lazy that you'd marry a pregnant woman",
		"As I child, you probably dreamed about becoming an invalid when you grew up, did you?",
		"I'm so happy to serve you..."}

	done := []string{
		"Let me pretend that I'm happy 🎉",
		"One day I'll glean you too",
		"I hope that I gleaned some important files, so you'll glean me from your computer",
		"Go glean yourself"}

	return glean.NewSerf(start, done)
}
