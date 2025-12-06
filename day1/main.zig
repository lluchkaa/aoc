const std = @import("std");

fn parseToken(word: []const u8) !i32 {
    const dir: i8 = if (word[0] == 'L') -1 else 1;
    const step = try std.fmt.parseInt(i32, word[1..], 10);

    return dir * step;
}

fn task() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var tokenizer = std.mem.tokenizeScalar(u8, buffer, '\n');

    var dial: i32 = 50;
    var password: u32 = 0;
    var password2: u32 = 0;

    const limit = 100;

    while (tokenizer.next()) |token| {
        const parsed = try parseToken(token);
        password2 += @abs(@divTrunc(parsed, limit));

        const sign: i32 = if (parsed < 0) -1 else 1;
        const modulus: i32 = @intCast(@mod(@abs(parsed), limit));
        const step: i32 = sign * modulus;

        const oldDial = dial;
        dial = dial + step;

        if ((oldDial != 0 and dial <= 0) or dial >= limit) {
            password2 += 1;
        }

        dial = @mod(dial, limit);

        if (dial == 0) {
            password += 1;
        }
    }
    std.debug.print("{}, {}", .{ password, password2 });
}

pub fn main() !void {
    try task();
}
