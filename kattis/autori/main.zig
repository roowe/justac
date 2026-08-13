const std = @import("std");

const Scanner = struct {
    tokens: std.mem.TokenIterator(u8, .any),

    fn init(input: []const u8) Scanner {
        return .{
            .tokens = std.mem.tokenizeAny(u8, input, "- \t\r\n"),
        };
    }
    fn nextString(self: *Scanner) ?[]const u8 {
        return self.tokens.next();
    }
};

pub fn main(init: std.process.Init) !void {
    var read_buffer: [4096]u8 = undefined;
    var stdin_reader =
        std.Io.File.stdin().readerStreaming(init.io, &read_buffer);

    const input = try stdin_reader.interface.allocRemaining(
        init.arena.allocator(),
        .limited(4096),
    );

    var scanner = Scanner.init(input);
    var output_buffer: [128]u8 = undefined;
    var n: usize = 0;
    while (scanner.nextString()) |word| {
        output_buffer[n] = word[0];
        n += 1;
    }
    output_buffer[n] = '\n';
    try std.Io.File.stdout().writeStreamingAll(init.io, output_buffer[0 .. n + 1]);
}
