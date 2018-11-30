# GoST

**GoST** is a *dumb* Go script that check constantly the *uptime*/*status code* of a WebServer.
GoST stands for **Go Site Tracker**

# Why?
Mainly because I'm bored. I don't have any reason to use a program like this. It's an exercise. I want to make practice with GoLang

# How?
**GoST** has no GUI. So it's a CLI program.
You need to pass arguments to gost to make it run.
## Arguments
 1) `<url>` is the target you want to track.
 2) `<request delay>` (in seconds) is the delay between checks.
 3) `<logfile>` (optional) is where GoST should pipe the output.
 
## Examples:
 - `gost https://foo.bar`
 - `gost https://jackrendor.cf 4 jackrendor_site.log`
 - `gost https://google.com 50 google.log`
## Output example:
### Command:
`gost https://jackrendor.cf 1 foobar.log`
### Output: 
```
2018-11-26 16:51:01 SUCCESS 200
2018-11-26 16:51:12 SUCCESS 200
2018-11-26 16:51:23 UNKOWN -1 Get https://jackrendor.cf: dial tcp 163.172.187.71:443: connect: network is unreachable
```

# Install
 1) First of all, [install go](https://golang.org/doc/install).
 2) Run `./build.sh` to create an executable inside the `bin/` folder.
 3) Run `sudo ./install.sh` or `su -c ./install.sh` to place the file inside `/usr/local/bin/`.
 
 # Thanks to:
 - Me [Telegram](https://t.me/jackrendor)
 - And Me [Linkem](https://it.linkedin.com/in/jackrendor)
 - Also me [GitLab](https://gitlab.com/jackrendor)
 - And the GoLang Italian Community [Telegram](https://t.me/golangit)

