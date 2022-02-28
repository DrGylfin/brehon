# brehon

Named after the Old Irish title for "judge," this is a to-do list with a random selector.

I realize that to-do list apps are a dime a dozen these days, but I wanted something a *tiny* bit differently.  Because I have ADHD, managing my various tasks, deadlines, and the like can be difficult.  But it's not just keeping track, it's also choosing; things can often seem equivalent.  For example, when I have free time, do I play game x or watch something on youtube?  So I wanted to have an app that'll let me organize what I have to do in ways that are helpful for me, which includes choosing randomly either from a specific category of tasks or from my whole list.

This is also a simple enough project from a programming standpoint that it seemed like a good starting place.

## Current functionality

- The UI sucks, but is at least functional on a basic level
- Supports the basic functions of a to-do list: adding and removing items
- File storage: can save the list to a file, loads from that list on start-up
	- Note: currently the data file is hard-coded, but this will change eventually
- Pick a task at random


## Tech Stack

This project is built using [Wails](https://wails.io/), a far better alternative to Electron.  It uses Go for the back-end, and doesn't pack a whole Chromium instance into every app.  I've found it to be a good balance between functionality and size, with less bloat than Electron.  It also does well for creating my first development project in awhile, since most of the functionality can be done in HTML, CSS, and JavaScript, with which I was most familiar.

## Credits & Thanks

I'd like to think the /r/webdev [discord](https://discord.gg/H9Jkc7p) for helping me with JavaScript, and the [#gophers](https://gophers.slack.com/messages/CJ4P9F7MZ/) Slack channel for helping me suck less at Go.

The code is all done by me, except for the Wails part.  

In-app icons are from [css.gg](https://css.gg).

Application icon is by [Elegant Themes](https://www.elegantthemes.com/blog/freebie-of-the-week/beautiful-flat-icons-for-free).