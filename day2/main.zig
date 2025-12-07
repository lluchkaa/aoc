const std = @import("std");

fn parseToken(word: []const u8) !struct { start: i64, end: i64 } {
    const trimmed = std.mem.trim(u8, word, " \n\r");
    const index = std.mem.indexOf(u8, trimmed, "-");

    const start = try std.fmt.parseInt(i64, trimmed[0..index.?], 10);
    const end = try std.fmt.parseInt(i64, trimmed[index.? + 1 ..], 10);

    return .{ .start = start, .end = end };
}

fn isInvalidID(id: i64) bool {
    var buf: [40]u8 = undefined;
    const width = std.fmt.printInt(buf[0..], id, 10, std.fmt.Case.lower, .{});

    for (1..@divTrunc(width, 2) + 1) |len| {
        if (@mod(width, len) == 0 and isInvalidIDForLen(buf[0..width], len)) {
            std.debug.print("Invalid {}\n", .{id});
            return true;
        }
    }

    return false;
}

fn isInvalidIDForLen(id: []u8, len: usize) bool {
    const parts: usize = @intCast(@divTrunc(id.len, len));
    const first = id[0..len];
    for (1..parts) |start| {
        if (!std.mem.eql(u8, first, id[start * len .. (start + 1) * len])) {
            return false;
        }
    }
    return true;
}

fn task() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var tokenizer = std.mem.tokenizeScalar(u8, buffer, ',');
    var sum: usize = 0;

    while (tokenizer.next()) |token| {
        const parsed = try parseToken(token);
        const start: usize = @intCast(parsed.start);
        const end: usize = @intCast(parsed.end);

        for (start..end) |num| {
            if (isInvalidID(@intCast(num))) {
                sum += num;
            }
        }
    }
    std.debug.print("Sum of invalid IDs: {d}\n", .{sum});
}

pub fn main() !void {
    try task();
}
