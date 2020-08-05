# xkcdbot

A basic discord bot written in Go that searches for xkcd because we all love
that don't we.

Made using [discordgo](https://github.com/bwmarrin/discordgo/)

[This repository](https://github.com/EdmundMartin/gosearcher) was helpful for me
in writing the backend for searches.

Hosted on [heroku](https://heroku.com/)

## Add to server

[invite link
here](https://discord.com/oauth2/authorize?client_id=738373714705514507&scope=bot&permissions=18432)

Permissions are to:

* send messages
* embed links

***OR***

You can clone the repository, and build the bot yourself. If you do this, you
can add your bot to your servers yourself :)

Bot token is stored in the environment variable `xkcdbottoken`.

```bash
git clone https://github.com/PioKozi/xkcdbot-discord.git
cd xkcdbot-discord

export xkcdbottoken="put your token here :)"

go build
./xkcdbot-discord
```

## Usage

Existing commands are:

* `.xkcd <string>` - searches Google with the string and returns the first
   result.  Also uses inurl: and site:, so as to avoid searching sites other
   than xkcd (on-topic and nsfw).
* `.xkcd` - returns a link to the most recent xkcd
* `.xkcdid <int>` - returns the link, directly to the xkcd with the id of the
   int given. This is faster than `.xkcd`.
* `.whatif <string>` - like `.xkcd`, but for what-if.xkcd.com
* `.whatifid <int>` - ditto but using id

Certain xkcds are also "cached", meaning if you search for it with a specific
term via `.xkcd <string>`, it will not be searched, but rather returned
*immediately* by id. An example of this is with `.xkcd security`, which
immediately posts [xkcd: Security](https://xkcd.com/538/)

## Todo

* [x] Figure out a FOSS hosting solution for this so I don't need to have this
    bot running on my laptop all the time.
* [x] Cache searches with an in-built type (lost on restart).
* [ ] Keep cache around, initialising it on restart, most likely using a
    relation DB.
* [ ] Migrate from searching xkcds via Google, to having all the xkcds in a
    postgresql DB, and make this work with heroku.
  * This should be useful, as currently the bot can be blocked by Google,
      which is bad for just wanting to use the bot :(
