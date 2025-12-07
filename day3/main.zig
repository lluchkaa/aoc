const std = @import("std");

fn getMaxJoltage(list: []const u8) i32 {
    const firstIndex = std.mem.indexOfMax(u8, list[0..(list.len - 1)]);
    const secondIndex = std.mem.indexOfMax(u8, list[(firstIndex + 1)..]);

    const firstDigit = list[firstIndex] - '0';
    const secondDigit = list[secondIndex + firstIndex + 1] - '0';
    return firstDigit * 10 + secondDigit;
}

fn task() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var tokenizer = std.mem.tokenizeScalar(u8, buffer, '\n');
    var sum: i32 = 0;

    while (tokenizer.next()) |token| {
        sum += getMaxJoltage(token);
    }
    std.debug.print("Total joltage: {d}\n", .{sum});
}

pub fn main() !void {
    try task();
}
