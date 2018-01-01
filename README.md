*Draft*

`glean` is a tool for photographers like myself who prefer taking photos in RAW *and* JPEG format.

## A simple usage case

My workflow is like follows:

1. review photos in JPEG (because it takes a lot of time to cycle through RAW files)
2. save the names and remove all the bad photos in JPEG
3. remove RAW files corresponding to the names of bad photos in JPEG

It's simple enough, but, on the other hand, it's kinda annoying that I have to manually write down names of the bad photos and then manually remove the RAW files.

`glean` is designed to do this for me automatically.

With `glean` the same workflow looks like this:

1. review photos in JPEG (because it takes a lot of time to cycle through RAW files)
2. remove all the bad photos
3. run `glean` and it will automatically remove corresponding RAW files

## Configuration

By default, `glean` is looking for RAW images in the same directory where the JPEGs are.
It also assumes that RAW files have the `.nef` extension which is Nikon's RAW format.

Chances are that there are people with a different camera make and / or different file organization strategy.
To change the default configuration, one has to create a configuration file which will instruct `glean` how to do its job.

If you run the `glean` command from a directory which has a file called `glean.yaml`, it will automatically use this file as a configuration.

Otherwise, you have to tell it where to look for this file like follows:
`glean --config=/Users/Santa/glean.conf`

### Configuration Syntax

`glean` expects configuration files to be in [YAML](https://en.wikipedia.org/wiki/YAML) format.

A full configuration file looks like this:

```
---
jpeg_dir: /Users/mobileka/photos/jpeg #use ./ for the current directory
jpeg_ext:
    - jpeg
    - jpg

raw_dir: /Users/mobileka/photos/raw #use ./ for the current directory
raw_ext:
    - nef
    - raw
    - dng
```
