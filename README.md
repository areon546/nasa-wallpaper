# Nasa Photo Background

## Disclaimer

This system is designed primarily for linux.
I don't know how to edit Windows Registry files,
however in theory that would be a much simpler solution to this.
It definitely won't work on windows.

Do note, if you are on Windows (10, maybe higher),
you can bypass restrictions on setting your background image pretty easily
and manually set registry file objects to bypass menu customisation options.

This project now works, to the best of my knowledge.
So I will now move onto another project since I don't have
the time to make the quality of this project better.
If someone is interested and encounters an error,
I will try my best to assist.

## Setup

The `Makefile` contains a collection of handy commands
that I used throughout making this process.

Run `make help` to get a description of what they do.

`make systemd` sets up the program and restarts the sytem services daemon
so that it recognises `nasa.service`, and then runs it.

## Project

The project taught me some basics on `systemd` services and timers,
implementation thereof, as well as using the `BurntSushi/toml`, and `encode/json` golang packages.

### Project Structure

`cmd` contains command line tools.
`internal` contains logic regarding processing data, and other bootstrapping of the project (miscellaneous setup)

I think a `domain` should be added, those usually contain core logic in the application space,
which in this case I expect would be the NASA api's.
Currently all of this is placed within `cmd/set-background` since it's only
1 API I am interfacing in.

This project so far is hard coded to support the `apod` [api](https://github.com/nasa/apod-api),
however I tried to design it so that it would be easier to restructure it
and add additional APIs (hence the arguably awkward [example configuration file](./example.toml)
.

Command Lines

- `set-background`: returns the URL of the configured APOD settings, and sets it as the user's background

## Current Dependencies

- `feh`, [feh](https://wiki.archlinux.org/title/Feh) command opens pictures and can set them as background
  - this works on I3WM (my current window manager), which uses [X11 as the Window Server](https://en.wikipedia.org/wiki/X_Window_System),
  I don't have a Hyperland system to test on
- make it so that more options are viable (I don't know which ones)
