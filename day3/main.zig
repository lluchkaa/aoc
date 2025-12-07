const std = @import("std");

const size: usize = 12;

fn getMaxJoltage(list: []const u8) i64 {
    var indexes: [size]usize = undefined;

    for (0..size) |i| {
        const start = if (i == 0) 0 else indexes[i - 1] + 1;
        const end = list.len - (size - i) + 1;

        const index = std.mem.indexOfMax(u8, list[start..end]);

        indexes[i] = index + start;
    }

    var sum: i64 = 0;

    for (indexes, 0..) |index, i| {
        sum += (list[index] - '0') * std.math.pow(i64, 10, @intCast(size - i - 1));
    }

    return sum;
}

fn task() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var tokenizer = std.mem.tokenizeScalar(u8, buffer, '\n');
    var sum: i64 = 0;

    while (tokenizer.next()) |token| {
        sum += getMaxJoltage(token);
    }
    std.debug.print("Total joltage: {d}\n", .{sum});
}

pub fn main() !void {
    try task();
}
