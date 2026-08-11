const std = @import("std");

pub fn main(init: std.process.Init) !void {
    const allocator = init.arena.allocator();

    var stdin_buffer: [4096]u8 = undefined;
    var stdin_reader = std.Io.File.stdin().reader(init.io, &stdin_buffer);
    const input = try stdin_reader.interface.allocRemaining(allocator, .unlimited);

    var lines = std.mem.tokenizeAny(u8, input, "\t\r\n");

    const count_line = lines.next() orelse return;
    const count = try std.fmt.parseInt(
        usize,
        count_line,
        10,
    );

    var stdout_buffer: [4096]u8 = undefined;
    var stdout_writer = std.Io.File.stdout().writer(init.io, &stdout_buffer);
    const stdout = &stdout_writer.interface;

    for (0..count) |i| {
        const raw_line = lines.next() orelse return error.MissingInput;
        //try stdout.print("i={d}\n", .{i});

        if (i % 2 == 0) {
            try stdout.print("{s}\n", .{raw_line});
        }
    }

    try stdout.flush();
}
