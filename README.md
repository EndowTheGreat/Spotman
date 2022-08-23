<h1 align=center>Spotman</h1>

**Spotman** is a command-line controller and metadata gathering tool for the Spotify media player on Linux.

----

<h2>Installation</h2>

```
$ git clone https://gitlab.com/EndowTheGreat/spotman.git
$ cd spotman
$ make
```

<h2>Usage</h2>

```
Usage:
  spotman [command]

Available Commands:
  help        Help about any command
  metadata    Fetch the metadata for the current song
  next        Skip to the next song
  previous    Skip to the previous song
  playpause   Toggle play/pause on the current song
  pause       Pause the current song
  play        Play the current song

Flags:
  -h, --help   help for spotman

Use "spotman [command] --help" for more information about a command.
```

Additionally, you can specify which metadata attribute to fetch.
```
Fetch the metadata for the current song

Usage:
  spotman metadata [flags]

Flags:
      --attribute string   Metadata attribute to fetch
  -h, --help               help for metadata
```

An example of this would look like this:
```
$ spotman metadata --attribute title
```

<h2>Contributing</h2>

Contributions are always welcome. Spotman is built with extensibility in mind, and it is relatively easy to create new commands. If you plan on doing so, a helpful resource for this is the [MPRIS](https://specifications.freedesktop.org/mpris-spec/latest/) D-Bus Interface Specification.
