const std = @import("std");

fn parseToken(word: []const u8) !struct { i8, i32 } {
    const dir: i8 = if (word[0] == 'L') -1 else 1;
    const step = try std.fmt.parseInt(i32, word[1..], 10);

    return .{ dir, step };
}

fn task1() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var tokenizer = std.mem.tokenizeScalar(u8, buffer, '\n');

    var start: i32 = 50;
    var password: i32 = 0;

    while (tokenizer.next()) |token| {
        const parsed = try parseToken(token);
        start += parsed.@"0" * parsed.@"1";
        start = @mod(start, 100);
        // std.debug.print("{}, {}", .{start, password});
        if (start == 0) {
            password += 1;
        }
    }
    std.debug.print("{}", .{password});
}

pub fn main() !void {
    try task1();
}
