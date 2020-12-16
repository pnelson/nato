# nato

## Usage

`nato` takes a single argument, the alphanumeric string to phoneticize.

```sh
$ nato aTczmP6S
alfa TANGO charlie zulu mike PAPA six SIERRA
```

Optionally use `-n` for newlines instead of spaces.

```sh
$ nato -n aTczmP6S
alfa
TANGO
charlie
zulu
mike
PAPA
six
SIERRA
```

## Motivation

I occasionally find myself speaking with a representative over the phone and
there is often some form of verification for which I've chosen a random set of
characters. I usually just read it out but I noticed that most often it is
repeated back to me with NATO phonetics. I see how it is useful but I don't
encounter this often enough to remember and so I built this tool to help me.
