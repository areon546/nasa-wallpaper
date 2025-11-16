# Nasa Photo Background 


Command Lines 

- `set-background`: returns the URL of the configured APOD settings, and sets it as the user's background



# Status:


Reads the following config file in the user's come directory: (`~` represents the home directory)

```
~/.config/dragon/nasa/wallpaper.toml
```

And uses it as a basis for the rest of the program.

The `{insert}` command runs the entire process

`make setup`: sets up the package to create a service, and add it to `systemd` to run repeatedly.


## Current Dependencies:

- `feh`, [feh](https://wiki.archlinux.org/title/Feh) command opens pictures and can set them as background 
- make it so that more are viable 












# Old
To Include:

- <https://github.com/BurntSushi/toml/tree/master> either through my files package, or directly, for TOML editing.
- <https://github.com/jpillora/opts> for command line arguments, seems useful
