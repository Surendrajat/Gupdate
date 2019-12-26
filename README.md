# GUpdate
Update current GO installation to the latest version available

## How?
- Clone repo and build `gupdate`

    `git clone https://github.com/surendrajat/gupdate && cd gupdate && go build .`
- Run `gupdate`

    `./gupdate` or `sudo ./gupdate` depending on your installation dir
## Why?
Frustration little, little laziness.
> "It is intended that programs written to the Go 1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification."

So why don't I use latest GO as soon it's released? Laziness, perhaps.
Downloading and extracting each time made me think why not write a shell script? But then again, GO is perfect for this task.
Also, why don't I just use some existing version manager or `go get ...`? I simply want to update the existing installation and nothing else.
