const std = @import("std");

const size: usize = 12;

const directions = [_][2]i8{
    .{ 0, -1 },
    .{ 1, -1 },
    .{ 1, 0 },
    .{ 1, 1 },
    .{ 0, 1 },
    .{ -1, 1 },
    .{ -1, 0 },
    .{ -1, -1 },
};

fn isAccessible(i: i32, j: i32, grid: [][]u8) bool {
    var count: i32 = 0;
    for (directions) |direction| {
        const x = j + direction[0];
        const y = i + direction[1];

        if (y >= 0 and x >= 0 and y < grid.len and x < grid[@intCast(y)].len and grid[@intCast(y)][@intCast(x)] == '@') {
            count += 1;
        }
    }
    return count < 4;
}

fn task() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const stat = try file.stat();

    const buffer = try file.readToEndAlloc(allocator, stat.size);

    var iterator = std.mem.splitScalar(u8, buffer, '\n');

    var grid = std.ArrayList([]u8).empty;
    defer grid.deinit(allocator);

    while (iterator.next()) |line| {
        const copy = try allocator.alloc(u8, line.len);
        std.mem.copyForwards(u8, copy, line);
        try grid.append(allocator, copy);
    }

    var count: i32 = 0;
    while (true) {
        var localCount: i32 = 0;
        for (grid.items, 0..) |row, i| {
            for (0..row.len) |j| {
                if (grid.items[i][j] == '@' and isAccessible(@intCast(i), @intCast(j), grid.items)) {
                    localCount += 1;
                    grid.items[i][j] = '.';
                }
            }
        }

        count += localCount;

        if (localCount == 0) {
            break;
        }
    }

    std.debug.print("Total count: {d}\n", .{count});
}

pub fn main() !void {
    try task();
}
