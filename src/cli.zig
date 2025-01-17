const std = @import("std");
const mem = std.mem;

const default_out = "./colorstorm-out";
const flag_prefix_long = "--";
const flag_prefix_short = "-";

const allocator = std.heap.page_allocator;
var settings = std.BufMap.init(allocator);

pub const Gen = enum { vim, vscode, sublime, atom, iterm2, all };
pub const Flag = enum { outdir, input, gen, help, na };

/// Initializes the CLI bufmap with values that are implicit unless
/// overwritten by the user.
pub fn init() !void {
    try settings.put(@tagName(Flag.outdir), default_out);
    try settings.put(@tagName(Flag.gen), @tagName(Gen.all));
    try settings.put(@tagName(Flag.input), "");
}

/// Parses a command line flag into a valid Flag enum type. Note that
/// the user can provide a long or short version of the flag. For example:
/// "--<flag_name>" or "-<letter>"
/// "--out" or "-o"
pub fn parse_flag(argument: []const u8) Flag {
    inline for (std.meta.fields(Flag)) |f| {
        if (std.mem.eql(u8, argument, flag_prefix_long ++ f.name) or
            std.mem.eql(u8, argument, flag_prefix_short ++ [_]u8{f.name[0]}))
        {
            return @enumFromInt(f.value);
        }
    }

    return Flag.na;
}

/// Sets or overwrites a flag's value using the user's specified value
pub fn set_flag_val(flag: Flag, argument: []const u8) !void {
    if (argument.len == 0) {
        return;
    }

    try settings.put(@tagName(flag), argument);
}

/// Retrieves a flag's default or overwritten value
pub fn get_flag_val(comptime flag: Flag) ?[]const u8 {
    return settings.get(@tagName(flag));
}

test "using long and short versions of flag" {
    const input_long: []const u8 = "--input";
    const input_short: []const u8 = "-i";
    const input_flag_long = parse_flag(input_long);
    const input_flag_short = parse_flag(input_short);

    try std.testing.expect(input_flag_long == input_flag_short);
}

test "check fallback flag value for invalid arguments" {
    const bad_input: []const u8 = "--invalid";
    const bad_flag = parse_flag(bad_input);

    try std.testing.expect(bad_flag == Flag.na);
}

test "set and get flag value" {
    const gen_flag: []const u8 = "vim";
    try set_flag_val(Flag.gen, gen_flag);

    try std.testing.expect(std.mem.eql(u8, get_flag_val(Flag.gen).?, gen_flag));
}
