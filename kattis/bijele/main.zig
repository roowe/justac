const std = @import("std");

pub fn main(init: std.process.Init) !void {
    const k: [6]i32 = .{ 1, 1, 2, 2, 2, 8 };
    var b: [6]i32 = .{ 0, 0, 0, 0, 0, 0 };
    var input_buffer: [64]u8 = undefined;
    var stdin_reader = std.Io.File.stdin().readerStreaming(init.io, &input_buffer);
    const line = try stdin_reader.interface.takeDelimiter('\n') orelse return error.MissingInput;
    var tokens = std.mem.tokenizeAny(u8, line, "\t\r ");
    var i: usize = 0;
    while (tokens.next()) |token| {
        b[i] = k[i] - try std.fmt.parseInt(i32, token, 10);
        i += 1;
    }

    var stdout_buffer: [32]u8 = undefined;
    const output = try std.fmt.bufPrint(&stdout_buffer, "{d} {d} {d} {d} {d} {d}\n", .{ b[0], b[1], b[2], b[3], b[4], b[5] });
    try std.Io.File.stdout().writeStreamingAll(init.io, output);
}
