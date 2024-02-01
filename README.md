# Qwertyflip

qwertyflip transforms input files by flipping their characters based on keyboard positions.

You can run it as:

```bash
./qwertyflip <command file> <input file>
```

where `command file` is a file containing a string of comma-separated transforms. The transforms can
be one of `H` for a horizontal keyboard flip, `V` for a vertical keyboard flip, or `+/-N` for a shift
by N positions. So the command `H,-3,V` would flip an input of `1` first horizontally to `0`, then
shift by -3 positions to `7`, then flip vertically to `m`.

Output is produced to standard out.
