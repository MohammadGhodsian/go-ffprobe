# go-ffprobe

This project is forked from https://github.com/vansante/go-ffprobe to handle multicast programs.

Small library for executing an ffprobe process on a given file and getting an easy to use struct
representing the returned ffprobe data.

## Installation

```
go get gopkg.in/MohammadGhodsian/go-ffprobe.v3
```

## Basic usage

To get a quick the quick data on a video file:

```golang
ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelFn()

data, err := ffprobe.ProbeURL(ctx, "/path/to/file.mp4")
if err != nil {
    log.Panicf("Error getting data: %v", err)
}
```

To get the ffprobe data for a video file that is accessible via HTTP, you can use the same
command, but with an HTTP URL.

To get the data of a file you have an `io.Reader` for, use:

```golang
ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelFn()

fileReader, err := os.Open("/path/to/file.mp4")
if err != nil {
    log.Panicf("Error opening test file: %v", err)
}

data, err := ffprobe.ProbeReader(ctx, fileReader)
if err != nil {
    log.Panicf("Error getting data: %v", err)
}
```