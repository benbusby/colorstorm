const std = @import("std");

const allocator = std.heap.page_allocator;

/// Returns a [3]u8 corresponding to RGB percentages
/// derived from a hex color string (i.e. "#ffffff")
pub fn hex_to_percent(hex_color: []const u8) []const u8 {
    var result_i = hex_to_dec(hex_color);
    var result_f = [3]f32{ 0.0, 0.0, 0.0 };

    for (result_i) |val, i| {
        result_f[i] = @intToFloat(f32, val) / 255;
    }

    const result_str = std.fmt.allocPrint(
        allocator,
        "{d:.7} {d:.7} {d:.7}",
        .{ result_f[0], result_f[1], result_f[2] },
    ) catch "0 0 0";

    return result_str;
}

/// Returns an [3]i32 corresponding to RGB int values
/// derived from a hex color string (i.e. "#ffffff")
pub fn hex_to_dec(hex_color: []const u8) [3]i32 {
    comptime var i = 0;
    var result = [3]i32{ 0, 0, 0 };

    for (result) |_, index| {
        // Split hex color into separate R/G/B components
        //std.debug.print("{d}\n", .{index});
        const string = std.fmt.allocPrint(
            allocator,
            "+0x{s}",
            .{hex_color[index * 2 + 1 .. index * 2 + 3]},
        ) catch "+0xFF";

        defer allocator.free(string);
        //std.debug.print("{s}!\n", .{string});
        result[index] = std.fmt.parseInt(i32, string, 0) catch 0;

        i += 1;
    }

    return result;
}

test "hex color to decimal" {
    var hex: []const u8 = "#ff00aa";
    var hex_dec = hex_to_dec(hex);

    try std.testing.expect(hex_dec[0] == 255);
    try std.testing.expect(hex_dec[1] == 0);
    try std.testing.expect(hex_dec[2] == 170);
}

test "hex color to percentage" {
    var hex: []const u8 = "#ffffff";
    var hex_percent = hex_to_percent(hex);
    var expected: []const u8 = "1.0000000";

    var iter = std.mem.split(hex_percent, " ");
    while (iter.next()) |c_val| {
        try std.testing.expect(std.mem.eql(u8, expected, c_val));
    }
}
