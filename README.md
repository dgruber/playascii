# playascii

playascii plays a local [asciinema](https://asciinema.org/) _cast_ file in a local browser
without uploading it somewhere else. It is _asciinema play_ for the local browser.

## How to use it

Record your terminal session with [asciinema](https://asciinema.org/docs/getting-started)
in a file.

    asciinema rec my.cast

Play the file in your browser.

    playascii my.cast


## 3rd Party Licenses

It embeds [asciinema-player](https://github.com/asciinema/asciinema-player/releases) in
player-css.go and player-js.go.

See https://github.com/asciinema/asciinema-player/blob/develop/LICENSE

For the remaining ones please check the vendor directory.