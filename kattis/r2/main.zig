const std = @import("std");

pub fn main(init: std.process.Init) !void {
    var input_buffer: [64]u8 = undefined;
    var stdin_reader = std.Io.File.stdin().readerStreaming(init.io, &input_buffer);
    const line = try stdin_reader.interface.takeDelimiter('\n') orelse return error.MissingInput;
    var tokens = std.mem.tokenizeAny(u8, line, "\t\r ");
    const r1 = try std.fmt.parseInt(i32, tokens.next().?, 10);
    const s = try std.fmt.parseInt(i32, tokens.next().?, 10);

    const r2 = 2 * s - r1;
    var stdout_buffer: [32]u8 = undefined;
    const output = try std.fmt.bufPrint(&stdout_buffer, "{d}\n", .{r2});
    try std.Io.File.stdout().writeStreamingAll(init.io, output);
}
