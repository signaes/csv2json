# csv2json

## Building

```
go build -o csv2json ./main.go
```

## Running

Let’s say we have a `csv` file with the following contents locally, named `sample.csv`:

```
a;b;c;d
1;2;3;4
```

If you then run:

```
./csv2json sample.csv
```

The following should be printed on your terminal:

```
[
  {
    "a": "1",
    "b": "2",
    "c": "3",
    "d": "4"
  }
]
```

If you move the `csv2json` binary to some directory in your `PATH` like `/usr/local/bin` for example, then you can have it available as a command on any terminal session.
