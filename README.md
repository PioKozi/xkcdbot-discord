# xkcdbot

## Add to server

[invite](https://discord.com/oauth2/authorize?client_id=738373714705514507&scope=bot&permissions=18432)

## Usage

Existing commands are:

* .xkcd \<string\> - searches Google with the string and returns the first result.
    Also uses inurl: and site:, so as to avoid searching sites other than xkcd
    (on-topic and avoid nsfw)
* .xkcdid \<int\> - returns the link, directly to the xkcd with the id of the int
    given. This is far faster than .xkcd.
* .whatif \<string\> - like .xkcd, but for what-if.xkcd.com
* .whatifid \<int\> - ditto but using id
