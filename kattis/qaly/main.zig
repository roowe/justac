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
    var total: f64 = 0.0;

    for (0..n) |_| {
        const quality = try scanner.nextFloat(f64);
        const years = try scanner.nextFloat(f64);

        total += quality * years;
    }

    var output_buffer: [64]u8 = undefined;
    const output = try std.fmt.bufPrint(
        &output_buffer,
        "{d}\n",
        .{total},
    );

    try std.Io.File.stdout().writeStreamingAll(init.io, output);
}
