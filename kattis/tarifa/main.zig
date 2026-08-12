const std = @import("std");

fn read_int(reader: *std.Io.Reader) !i32 {
    const line = try reader.takeDelimiter('\n') orelse return error.EndOfStream;
    return try std.fmt.parseInt(i32, std.mem.trim(u8, line, "\t\r "), 10);
}

pub fn main(init: std.process.Init) !void {
    var input_buffer: [64]u8 = undefined;
    var stdin_reader = std.Io.File.stdin().readerStreaming(init.io, &input_buffer);
    const x = try read_int(&stdin_reader.interface);
    const n = try read_int(&stdin_reader.interface);
    var s: i32 = (n + 1) * x;

    for (0..@intCast(n)) |_| {
        const p = try read_int(&stdin_reader.interface);
        s -= p;
    }

    var stdout_buffer: [32]u8 = undefined;
    const output = try std.fmt.bufPrint(&stdout_buffer, "{d}\n", .{s});
    try std.Io.File.stdout().writeStreamingAll(init.io, output);
}
