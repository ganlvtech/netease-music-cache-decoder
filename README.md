# Netease Cloud Music Cache Decoder

XOR each byte with `0xa3`.

## Usage

Open `%LocalAppData%\Netease\CloudMusic\Cache\Cache\`.

Find the `.uc` file based on file size and last modified time.

```cmd
main.exe xxxxxxxxxx-xxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.uc
```

## Build

```bash
go build main.go
```

## License


               DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                       Version 2, December 2004

    Copyright (C) 2019 Ganlv

    Everyone is permitted to copy and distribute verbatim or modified
    copies of this license document, and changing it is allowed as long
    as the name is changed.

               DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
      TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

     0. You just DO WHAT THE FUCK YOU WANT TO.
