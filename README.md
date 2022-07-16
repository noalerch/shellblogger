# Shellblogger
A command-line blogging tool for managing posts on Hugo sites and deploying these to remote servers.

Work in progress!

## Usage
Get the application with `git clone https://github.com/noalerch/shellblogger`. Generate a binary file `shellblogger` with `go build`, or just run the process directly
with `go run main.go [COMMAND]`. You can run the binary with `./shellblogger` or add it to your PATH as `shblog` (pronounced like a single word).
Make a new Hugo website in your directory of choice. The easiest (and only possible so far) is to make it in the same directory as you installed shellblogger to.
When you have made and published a post, build your site directly with `shblog build`. Configure the program to go to your server of choice, then deploy
your changes with `shblog deploy`.

I have not yet made proper configuration of the application possible, so you will have to play around with the source code to push your site to your remote host.

## Dependencies
The program has the following dependencies:
- Go
- Hugo

## TODO
It's very early in the development process. I intend to implement the following:
- Quick posting
- Automatic Hugo site generation
- Configuration handling
- Better error handling
- Simple setup for Uppsala University students to generate websites at their user.it.uu.se adress.

