*Draft*

`glean` is a tool for photographers who prefer taking photos in RAW *and* JPEG at the same time.

## A simple usage case

My workflow is like follows:

1. review photos in JPEG (because it takes a lot of time to cycle through RAW files)
2. remove all bad photos in JPEG and save the names of the removed files
3. remove RAW files corresponding to the names of the bad photos in JPEG

It's simple enough, but, on the other hand, it's a bit annoying that I have to manually write down names of the bad JPEG photos and then manually remove the RAW files.

I built `glean` to do this for me automatically.

With `glean` the same workflow looks like this:

1. create a configuration file for `glean` (only once)
2. review photos in JPEG
3. remove all the bad JPEG photos
4. run `glean` and it will automatically remove corresponding RAW files

## Configuration

In order to minimize chances of accidentally removing valuable RAW files, `glean` requires you to create a configuration file.

If you run `glean` from a directory which has a file called `glean.yaml`, it will automatically use this file as a configuration.

Otherwise, you have to tell it where to look for this file like follows:
`glean --config=../relative/path/to/glean.conf`

### Configuration Syntax

`glean` expects configuration files to be in [YAML](https://en.wikipedia.org/wiki/YAML) format.

A configuration file should look like this:

```
---
jpeg_dir: . # should be relative to a directory from where you run `glean`. Confusing? I know. `.` sing means "current directory"
jpeg_ext:
    - .jpeg
    - .jpg

raw_dir: . # should be relative to a directory from where you run `glean`. Confusing? I know. `.` sing means "current directory"
raw_ext:
    - .nef
    - .raw
    - .dng
```

`jpeg_dir` - where are your reference JPEG files located
`jpeg_ext` - a case-sensitive list of "reference" file extensions

`raw_dir` - where to look for corresponding RAW files
`raw_ext` - a case-sensitive list of "victim" extensions

## Other usage scenarios

`glean` can be used with other file extensions provided that:

1. you have reference files with a certain extensions (e.g. `.ref`, and `.reference`)
2. you have other files with different extensions which have to be removed if there're no `.ref` files with the same name

# WARNING

`glean` does not ask for confirmations and it does not put anything into your "Trash".
If you made a mistake and it removed pictures of your 1 yo old daughter, please don't be upset with me.
