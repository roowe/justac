dir := invocation_directory()

build:
    cd {{ quote(dir) }} && zig build-exe -O ReleaseFast -femit-bin=main main.zig

run:
    cd {{ quote(dir) }} && zig run -O ReleaseFast main.zig

test input: build
    cd {{ quote(dir) }} && ./main < {{ quote(input) }}

alias b := build
alias r := run
alias t := test
