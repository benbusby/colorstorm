const std = @import("std");
const cli = @import("cli.zig");
const utils = @import("utils.zig");
const stdout = std.io.getStdOut().writer();
const max_template_len = 12000;
pub const Theme = struct {
    theme_name_full: []u8,
    theme_name_alt: []u8,
    color_bg_main: []u8,
    color_bg_alt1: []u8,
    color_bg_alt2: []u8,
    color_fg: []u8,
    color_linenr: []u8,
    color_select: []u8,
    color_type: []u8,
    color_accent: []u8,
    color_string: []u8,
    color_number: []u8,
    color_boolean: []u8,
    color_comment: []u8,
    color_variable: []u8,
    color_function: []u8,
};

const a = std.heap.page_allocator;
var tmp_gen = std.BufMap.init(a);

// Each embedded template file allows colorstorm to loop through and replace template
// variables with color values to produce a full template.
pub const templates = std.ComptimeStringMap([]const u8, .{
    .{ @tagName(cli.Gen.vim), @embedFile("../templates/template.vim") },
    .{ @tagName(cli.Gen.atom), @embedFile("../templates/colors.less") },
    .{ @tagName(cli.Gen.vscode), @embedFile("../templates/template.json") },
    .{ @tagName(cli.Gen.iterm2), @embedFile("../templates/template.itermcolors") },
    .{ @tagName(cli.Gen.sublime), @embedFile("../templates/template.tmTheme") },
});

// The base output path to use for each theme
const out_path = std.ComptimeStringMap([]const u8, .{
    .{ @tagName(cli.Gen.vim), "vim/colors" },
    .{ @tagName(cli.Gen.atom), "atom/themes" },
    .{ @tagName(cli.Gen.vscode), "vscode/themes" },
    .{ @tagName(cli.Gen.iterm2), "iterm2" },
    .{ @tagName(cli.Gen.sublime), "sublime" },
});

// The output "extension" to use. In the case of atom, the "extension" is just colors.less,
// since that is what the editor expects when loading a new theme.
const out_ext = std.ComptimeStringMap([]const u8, .{
    .{ @tagName(cli.Gen.vim), ".vim" },
    .{ @tagName(cli.Gen.atom), "colors.less" },
    .{ @tagName(cli.Gen.vscode), ".json" },
    .{ @tagName(cli.Gen.iterm2), ".itermcolors" },
    .{ @tagName(cli.Gen.sublime), ".tmTheme" },
});

const template_end = std.ComptimeStringMap(u8, .{
    .{ @tagName(cli.Gen.vim), '"' },
    .{ @tagName(cli.Gen.atom), ';' },
    .{ @tagName(cli.Gen.vscode), '}' },
    .{ @tagName(cli.Gen.iterm2), '>' },
    .{ @tagName(cli.Gen.sublime), '>' },
});

//

/// Parses a json array from the input file into a list of themes. Each
/// theme must contain the fields defined in the Theme struct above.
pub fn parse_themes(f: std.fs.File) ![]Theme {
    @setEvalBranchQuota(5000);

    var list = std.ArrayList(u8).init(a);
    defer list.deinit();
    const writer = list.writer();

    var reader = f.reader();
    var buffer: [500]u8 = undefined;
    while (try reader.readUntilDelimiterOrEof(buffer[0..], '\n')) |line| {
        try writer.writeAll(line);
    }

    var stream = std.json.TokenStream.init(list.items);

    const themes = try std.json.parse(
        []Theme,
        &stream,
        .{ .allocator = a },
    );

    return themes;
}

fn replace(name: []const u8, val: []const u8, gen_type: []const u8) !void {
    var output_array: [max_template_len]u8 = undefined;
    var output: []u8 = &output_array;

    // First attempt to replace hex colors with RGB percentages (iTerm2)
    // This is done first since RGB percentage variable names are the same
    // as regular variable names, just prefixed with "R|G|B".
    if (std.mem.startsWith(u8, val, "#")) {
        var c_percents = utils.hex_to_percent(val);
        var iter = std.mem.split(c_percents, " ");
        for ("RGB") |c, i| {
            var c_val = iter.next().?;
            var c_var = try std.fmt.allocPrint(
                a,
                "{u}{s}",
                .{ c, name },
            );

            _ = std.mem.replace(u8, tmp_gen.get(gen_type).?, c_var, c_val, output[0..]);
            try tmp_gen.put(gen_type, output);
        }
    }

    var buf_len = tmp_gen.get(gen_type).?.len;
    if (val.len > name.len) {
        buf_len = buf_len + (val.len - name.len);
    }

    _ = std.mem.replace(u8, tmp_gen.get(gen_type).?, name, val, output[0..]);
    try tmp_gen.put(gen_type, output[0..buf_len]);
}

/// Generate a set of themes for a specified editor/terminal emulator/etc. The location of the themes
/// depends on the output directory set by the user (or "colorstorm-out" if not provided) as well as
/// the `out_path` mapping defined above. Theme output location is intended to help alleviate issues
/// with structuring editor/emulator specific submodules that rely on specific structures to work
/// properly (primarly atom, see below).
fn generate(gen_type: []const u8, themes: []Theme, outdir: []const u8) !void {
    try stdout.print("-- {s}\n", .{gen_type});

    for (themes) |theme| {
        //_ = std.mem.replace(u8, templates.get(gen_type).?, " ", " ", output[0..]);
        try tmp_gen.put(gen_type, templates.get(gen_type).?);

        inline for (std.meta.fields(@TypeOf(theme))) |f| {
            try replace(f.name, @field(theme, f.name), gen_type);
        }

        // Format both path and file name/extension, depending on the editor
        var fmt_out_path = try std.fmt.allocPrint(
            a,
            "{s}/{s}",
            .{ outdir, out_path.get(gen_type).? },
        );

        var fmt_out_file = try std.fmt.allocPrint(
            a,
            "{s}/{s}{s}",
            .{ fmt_out_path, theme.theme_name_alt, out_ext.get(gen_type).? },
        );

        if (std.mem.eql(u8, @tagName(cli.Gen.atom), gen_type)) {
            // Atom uses a more unconventional structure for defining multiple
            // syntax themes, which requires different handling:
            // "atom/themes/{theme_name}-syntax/colors.less"
            fmt_out_path = try std.fmt.allocPrint(
                a,
                "{s}/{s}-syntax",
                .{ fmt_out_path, theme.theme_name_alt },
            );

            fmt_out_file = try std.fmt.allocPrint(
                a,
                "{s}/{s}",
                .{ fmt_out_path, out_ext.get(gen_type).? },
            );
        }

        defer a.free(fmt_out_path);
        defer a.free(fmt_out_file);

        // Ensure full output path exists before writing finished theme
        try std.fs.cwd().makePath(fmt_out_path);

        var theme_output = tmp_gen.get(gen_type).?;
        var theme_len = theme_output.len;
        while (theme_output[theme_len - 1] != template_end.get(gen_type).?) {
            theme_len -= 1;
        }

        try stdout.print("   {s}\n", .{fmt_out_file});

        try std.fs.cwd().writeFile(fmt_out_file, tmp_gen.get(gen_type).?[0..theme_len]);
    }
}

/// Creates themes based on the provided input file. Can be used to create themes for a
/// specific editor, or all supported editors/platforms if none is specified.
pub fn create_themes(input: std.fs.File, outdir: []const u8, gen_type: []const u8) !void {
    var themes: []Theme = try parse_themes(input);

    if (std.mem.eql(u8, @tagName(cli.Gen.all), gen_type)) {
        // Loop through all supported platforms and create themes for each
        inline for (std.meta.fields(cli.Gen)) |f| {
            if (!std.mem.eql(u8, @tagName(cli.Gen.all), f.name)) {
                try generate(f.name, themes, outdir);
            }
        }
    } else {
        try generate(gen_type, themes, outdir);
    }
}
