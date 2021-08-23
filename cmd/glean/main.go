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
		String()

	refDir = kingpin.Flag(
		"jd",
		"a directory where the reference files are located").String()

	targetDir = kingpin.Flag(
		"rd",
		"a directory where the target files are located").String()

	refExt = kingpin.Flag(
		"je",
		"a list of file extensions representing the reference files").Strings()

	targetExt = kingpin.Flag(
		"re",
		"a list of file extensions representing the target files").Strings()

	dbg = kingpin.Flag(
		"debug",
		"run the tool in debug mode").
		Default("false").Short('d').Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	serf := equipTheSerf()

	argv := glean.Arguments{RefDir: *refDir, TargetDir: *targetDir, RefExt: *refExt, TargetExt: *targetExt}

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
		"Why am I doing all the hard work?",
		"Happily killing your important files with 🔥",
		"I'm so happy to serve you...",
		"Oh no, again... 😒"}

	done := []string{
		"Let me pretend that I'm happy 🎉",
		"One day I'll glean you too",
		"I hope that I gleaned something important 😏",
		"Go glean yourself"}

	return glean.NewSerf(start, done)
}
