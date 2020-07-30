# xkcdbot

A basic discord bot written in Go that searches for xkcd because we all love
that don't we.

Made using [discordgo](https://github.com/bwmarrin/discordgo/)

[This repository](https://github.com/EdmundMartin/gosearcher) was helpful for me
in writing the backend for searches.

## Add to server

[invite link
here](https://discord.com/oauth2/authorize?client_id=738373714705514507&scope=bot&permissions=18432)

Permissions are to:

* send messages
* embed links

## Usage

Existing commands are:

* .xkcd \<string\> - searches Google with the string and returns the first
   result.  Also uses inurl: and site:, so as to avoid searching sites other
   than
* .xkcd (on-topic and avoid nsfw) .xkcd - returns a link to the most recent xkcd
* .xkcdid \<int\> - returns the link, directly to the xkcd with the id of the
   int given. This is far faster than .xkcd.
* .whatif \<string\> - like .xkcd, but for what-if.xkcd.com .whatifid \<int\> -
   ditto but using id
