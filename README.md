# Zeit

Zeit is an extremely simple command line application to read a natural language time description and output it as a timestamp.
The main usecase for this for me is quickly generating timestamps for discords timestamp values.

## Usage

`zeit yesterday`
`zeit -c -d three years ago 12pm`


## Flags

`-c`: automatically copy the generated output into your clipboard
`-d`: generate the timestamp as a discord timestamp value. This can then be pasted into discord, and will be displayed as the correctly timezone adjusted time for everyone.


## Why go?

This is practically just a more elaborate shellscript - go has the libraries I need.
Please don't see this as encouragement to use go for any real projects.

## Installation

Make sure you have go properly set up.
Then just run `go install github.com/elkowar/zeit` to install zeit!
