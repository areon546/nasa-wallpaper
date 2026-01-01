# NASA Photo Background

Check out [today's photo](https://apod.nasa.gov/apod/).

[Backgrounds I like](./cool-backgrounds.md)


## Disclaimer

This system is designed primarily for Linux.
I assume the equivalent solution for this on windows is editing 
Windows Registry file for the background, and then use the Windows Scheduler to set up 
a schedule for a similar time, and after bootup. 

Do note, if you are on Windows (10 tested, likely higher aswell),
you can bypass restrictions on setting your background image pretty easily
and manually set registry file objects to bypass menu customisation options.

If someone is interested and encounters an error,
I will do my best to assist.
If I get bugs in the future, I will fix as available.
If you use another background command, 
or know of a more standardised way to do this, 
assume I don't know about it and am interested in learning about it.


## Setup

The `Makefile` contains a collection of handy commands
that I used throughout making this process.

Run `make help` to get a description of what they do.

`make systemd` sets up the program and restarts the system services daemon
so that it recognises `nasa.service` and `nasa.timer`, and then runs it.

Steps:

1. Clone the repository.
2. Run `make systemd`

I am using systemd as my current init system, 
however that is no restriction.

## Improvements

- Implement windows version (somehow)?
- Make it more reliable (timer file needs to be edited as some days it doesn't seem to work)


## Project

The project taught me some basics on `systemd` services and timers,
implementation thereof, as well as using the `BurntSushi/toml`, and `encode/json` golang packages.

### Structure

- `cmd` contains command line tools.
- `internal` contains logic regarding processing data, and other bootstrapping of the project (miscellaneous setup)
- `pkg` contains important logic that I don't mind being public
- `domain` contains NASA api processing logic, not sure if it is appropriate however as it is using internal logic, 
which seems counter to what a good domain package would do

This project so far is hard coded to support the `apod` [api](https://github.com/nasa/apod-api),
however I tried to design it so that it would be easy to restructure it
and add additional APIs (hence the arguably awkward [example configuration file](./example.toml)),
however in theory so long as an API contains a `url` parameter in the json, like below,

```json
{
  ...
  "url":"",
  ...
}
```
it should function.

### CMD Commands

Command Lines

- `set-background`: returns the URL of the configured APOD settings, and sets it as the user's background

## Current Dependencies

- `feh`, [feh](https://wiki.archlinux.org/title/Feh) command opens pictures and can set them as background
  - this works on I3WM (my current window manager), which uses [X11 as the Window Server](https://en.wikipedia.org/wiki/X_Window_System),
  I don't have a Hyperland system to test on
- make it so that more options are viable (I don't know which ones)
