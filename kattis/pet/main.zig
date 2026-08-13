const std = @import("std");

const Scanner = struct {
    tokens: std.mem.TokenIterator(u8, .any),

    fn init(input: []const u8) Scanner {
        return .{
            .tokens = std.mem.tokenizeAny(u8, input, " \t\r\n"),
        };
    }

    fn nextInt(self: *Scanner, comptime T: type) !T {
        const token = self.tokens.next() orelse
            return error.MissingInput;

        return std.fmt.parseInt(T, token, 10);
    }

    fn nextFloat(self: *Scanner, comptime T: type) !T {
        const token = self.tokens.next() orelse
            return error.MissingInput;

        return std.fmt.parseFloat(T, token);
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
    var max_score: u64 = 0;
    var max_index: usize = 0;
    for (0..5) |i| {
        var score: u64 = 0;
        for (0..4) |_| {
            score += try scanner.nextInt(u64);
        }
        if (score > max_score) {
            max_score = score;
            max_index = i;
        }
    }

    var output_buffer: [64]u8 = undefined;
    const output = try std.fmt.bufPrint(&output_buffer, "{d} {d}\n", .{ max_index + 1, max_score });
    try std.Io.File.stdout().writeStreamingAll(init.io, output);
}
