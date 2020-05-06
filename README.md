# discotime

Discotime is the project that runs `tzbot`, a bot on discord that is intended to help users who schedule events and have attendees from multiple timezones.

## What can the bot do?

Right now, not much. But the planned features as follows:

## First, set a more accurate timezone for yourself

Timezones are complicated by the fact that some timezones change twice a year; for example "Pacific Standard Time" or PST changes from UTC-7 to UTC-8 and back again during the summer.

If you message the bot with the `!set` command you can set a canonical timezone, such as `America/Los_Angeles` as your timezone, and the bot will use that, instead of trying to guess based on your nickname in the channel.

You can view a list of canonical timezones at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

```
>>> !set America/Los_Angeles
tzbot: ok, your timezone is now set to America/Los_Angeles.
tzbot: I will assume this is your local timezone when I convert a time to your local time.
```

## Automatically guess your timezone based on your channel nickname.

If for example, you have a nickname in a server with `(-7)` or `[-7]` at the end of your name, `tzbot` will assume that's the time you want to convert to if you have not configured a more precise timezone using `!set`. This will allow you to (inside a channel only) do:

```
>>>>>> !localtime 5:54pm
tzbot: That time is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
tzbot: Local time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
```

This allows people who have yet to set a timezone but have it in their nickname to take advantage of the bot.

## Convert time from one timezone to another.

For example

```
>>>>>> !localtime 5:36pm in America/Los_Angeles
tzbot: That time is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
tzbot: Local time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
```

This will assume the time given is in UTC, and convert to America/Los_Angeles time (currently -0700 but outside of summer time it is -0800).

You can use timezone abbreviations:

```
>>>>>> !localtime 5:36pm in PDT
tzbot: That time is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
tzbot: Local time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
```

You can find a list of supported abbreviations here: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

Please note that these abbreviations have limitations. PST could mean "Pacific Standard Time" or "Philippines Standard Time", so I recommend you use the ones that are more specific.

You can use offsets:

```
>>>>>> !localtime 5:36pm in -0700
tzbot: That time is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
tzbot: Local time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
```

You can also convert from one timezone to another

```
>>>>>> !convert 8:54am PST to UTC
tzbot: That time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
tzbot: Which is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
```

If you want to remove ambiguity, then use the canonical timezone names from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

```
>>> !convert 10:36am America/Los_Angeles to UTC
tzbot: That time is 8:54am on Tuesday, 05 May 2020 PDT (UTC-0700)
tzbot: Which is 3:54pm on Tuesday, 05 May 2020 UTC (UTC+0000)
```

## Remind me

You can ask the bot to private message you at a specific time in UTC. If you give it no offset or timzeone it will use the configured timezone that you set earlier, and failing that it will try to use guess based on your in-server nickname:

```
>>> !remindme tomorrow at 11:30pm
tzbot: That time is 11:30pm on Wednesday, 06 May 2020 UTC (UTC+0000)
tzbot: I will remind you at 4:30pm on Wednesday, 06 May 2020 PDT (UTC-0700)
```

if the bot cannot infer your timezone, will will error.

You can also provide a timezone to the bot when asking for a reminder, in which case it will do the necessary conversion from the one given to the one you are in:

```
>>>>> !remindme tomorrow at 11:30pm +0800
tzbot: That time is 11:30pm on Wednesday, 06 May 2020 +0800 (UTC+0800)
tzbit: I will remind you at 8:30am on Wednesday, 06 May 2020 PDT (UTC-0700)
```

## Natural Language Dates

In order to try and be more user friendly, we try to parse the date/time provided using https://github.com/tj/go-naturaldate

For example, the following should be possible:

* now
* today
* yesterday
* 5 minutes ago
* three days ago
* last month
* next month
* one year from now
* yesterday at 10am
* last sunday at 5:30pm
* sunday at 22:45
* next January
* last February
* December 25th at 7:30am
* 10am
* 10:05pm
* 10:05:22pm
* 5 days from now
* 25th of December at 7:30am

# Adding the bot to my server

If you want this bot in your Discord server, please click [here](https://discord.com/api/oauth2/authorize?client_id=707063041547829279&permissions=388160&scope=bot).

Note that the bot is under active development.

# About hosting and data storage, privacy, etc

The bot is hosted in Heroku under a free tier account, with a redis database for persistant storage of timezone settings. You can see exactly the code that is pushed to Heroku, here in the repository. Note that I do log messages that the bot recieves for debugging purposes - this is all messages prefixed with a `!` - however Heroku [only keeps 1500 of the latest log messages](https://devcenter.heroku.com/articles/logging#log-history-limits) and given that I only log stuff interesting to the bot (potential commands) it's unlikely I'm going to see anything interesting. If I did, I have better things to do with my time than snoop people's private conversations.

# Contributing

If you want to help develop the bot, or extend it's capabilities, PRs are gratefully requested. If you don't have any coding ability, then feel free to create suggestions in Github's issue tracker for this project.

# License

```
MIT License

Copyright (c) 2020 Nathan Ollerenshaw

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
