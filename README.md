# InboxGenie
Emails on telegram

## Why? 

Sometimes my email app just fails at sending me notifications whenever I get an email. This
has caused me to miss urgent emails, which lead me to refresh my email app every 5 minutes. 

How ever, telegram notifications never fail, so I figured I should make a simple bot to send me
a message each whenever I get a new email on my inbox. 

## How? 

If for some reason you want to use it you'll need a few things: 

- Telegram bot API key
- Email service compatible with IMAP 
- Know your personal Telegram ID

Once you have all that you need to create a `config.env` like this: 

``` 
SERVER=imap.example.com:993
EMAIL=youremail@email.com
PASSWORD=Yourpassword
TGID=PersonalTelegramID
TGAPI=BotTelegramAPIKey
```

Once you do that you'll be able to run the bot locally

## Next?

I don't have defined roadmap yet, but some of the things I'd like to do are: 
- Get the full text of the email's body (must check for diffrent content-types)
- Being able to respond to emails by replying a message
- Build a sort of inbox functionallity. 

