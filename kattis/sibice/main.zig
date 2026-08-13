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

    const n = try scanner.nextInt(usize);
    const w = try scanner.nextInt(i32);
    const h = try scanner.nextInt(i32);
    for (0..n) |_| {
        const match = try scanner.nextInt(i32);
        if (match <= w or match <= h or match <= @sqrt(@as(f64, @floatFromInt(w * w + h * h)))) {
            try std.Io.File.stdout().writeStreamingAll(init.io, "DA\n");
        } else {
            try std.Io.File.stdout().writeStreamingAll(init.io, "NE\n");
        }
    }
}
