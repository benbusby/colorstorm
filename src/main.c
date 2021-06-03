#include <stdlib.h>
#include <stdio.h>
#include "themes.h"

Editor editors[N_EDITORS] = {
    {"templates/template.vim", "vim/colors/%s.vim"},
    {"templates/template.json", "vscode/themes/%s.json"},
    {"templates/earthbound_template.tmTheme", "sublime/earthbound_%s.tmTheme"},
    {"templates/colors.less", "atom/themes/%s-syntax/colors.less"}
};

Theme themes[N_THEMES] = {
    {
        "earthbound", 1, {
            {"theme_name_full", "Earthbound"},
            {"theme_name_alt",  "earthbound"},
            {"color_bg_main",   "#360a5f"},
            {"color_bg_alt1",   "#2b044f"},
            {"color_bg_alt2",   "#1c0037"},
            {"color_fg",        "#ffffff"},
            {"color_linenr",    "#94b2b2"},
            {"color_select",    "#6a6c23"},
            {"color_type",      "#f56b3c"},
            {"color_accent",    "#fafd51"},
            {"color_string",    "#84fc60"},
            {"color_boolean",   "#70caff"},
            {"color_variable",  "#abdcdc"},
            {"color_number",    "#fb967f"},
            {"color_comment",   "#acaf6f"},
            {"color_function",  "#70caff"}
        }
    },
    {
        "moonside", 0, {
            {"theme_name_full", "Moonside"},
            {"theme_name_alt",  "moonside"},
            {"color_bg_main",   "#000000"},
            {"color_bg_alt1",   "#080808"},
            {"color_bg_alt2",   "#131313"},
            {"color_fg",        "#ffffff"},
            {"color_linenr",    "#9e5dc8"},
            {"color_select",    "#5a1359"},
            {"color_type",      "#f6f929"},
            {"color_accent",    "#fd35fa"},
            {"color_string",    "#ff6693"},
            {"color_boolean",   "#fd9935"},
            {"color_variable",  "#c67ff4"},
            {"color_number",    "#aaef64"},
            {"color_comment",   "#7ca454"},
            {"color_function",  "#5e9aff"}
        }
    },
    {
        "threed", 1, {
            {"theme_name_full", "Threed"},
            {"theme_name_alt",  "threed"},
            {"color_bg_main",   "#303454"},
            {"color_bg_alt1",   "#373c60"},
            {"color_bg_alt2",   "#202339"},
            {"color_fg",        "#f0faff"},
            {"color_linenr",    "#9590b2"},
            {"color_select",    "#534f63"},
            {"color_type",      "#ffcfcb"},
            {"color_accent",    "#d4cbff"},
            {"color_string",    "#ffcf32"},
            {"color_boolean",   "#c67ff4"},
            {"color_variable",  "#2fff89"},
            {"color_number",    "#d4cbff"},
            {"color_comment",   "#bdb7db"},
            {"color_function",  "#f89070"}
        }
    },
    {
        "fire-spring", 1, {
            {"theme_name_full", "Fire Spring"},
            {"theme_name_alt",  "fire-spring"},
            {"color_bg_main",   "#261933"},
            {"color_bg_alt1",   "#21162c"},
            {"color_bg_alt2",   "#181020"},
            {"color_fg",        "#ffffca"},
            {"color_linenr",    "#b49a19"},
            {"color_select",    "#632611"},
            {"color_type",      "#ff7e50"},
            {"color_accent",    "#f0e500"},
            {"color_string",    "#74e4f3"},
            {"color_boolean",   "#d9c400"},
            {"color_variable",  "#e5caff"},
            {"color_number",    "#a99ade"},
            {"color_comment",   "#bb8673"},
            {"color_function",  "#d992ff"}
        }
    },
    {
        "dusty-dunes", 1, {
            {"theme_name_full", "Dusty Dunes"},
            {"theme_name_alt",  "dusty-dunes"},
            {"color_bg_main",   "#1e1b07"},
            {"color_bg_alt1",   "#150d00"},
            {"color_bg_alt2",   "#140f00"},
            {"color_fg",        "#f9e4a1"},
            {"color_linenr",    "#f9e4a1"},
            {"color_select",    "#6b5e33"},
            {"color_type",      "#e0c364"},
            {"color_accent",    "#f6d56a"},
            {"color_string",    "#ffebae"},
            {"color_boolean",   "#ffd03c"},
            {"color_variable",  "#f6d56a"},
            {"color_number",    "#f6d56a"},
            {"color_comment",   "#aaaa88"},
            {"color_function",  "#f6d56a"}
        }
    },
    {
        "magicant", 0, {
            {"theme_name_full", "Magicant (Light)"},
            {"theme_name_alt",  "magicant"},
            {"color_bg_main",   "#e6e5ab"},
            {"color_bg_alt1",   "#efeeb2"},
            {"color_bg_alt2",   "#f9f8b9"},
            {"color_fg",        "#220088"},
            {"color_linenr",    "#876a55"},
            {"color_select",    "#b299ff"},
            {"color_type",      "#881200"},
            {"color_accent",    "#604633"},
            {"color_string",    "#7100b1"},
            {"color_boolean",   "#a31100"},
            {"color_variable",  "#87000d"},
            {"color_number",    "#604633"},
            {"color_comment",   "#525252"},
            {"color_function",  "#9d02f2"}
        }
    },
    {
        "cave-of-the-past", 0, {
            {"theme_name_full", "Cave of the Past (Monochrome)"},
            {"theme_name_alt",  "cave-of-the-past"},
            {"color_bg_main",   "#b0d0b8"},
            {"color_bg_alt1",   "#a5c4ad"},
            {"color_bg_alt2",   "#9ab5a2"},
            {"color_fg",        "#262e25"},
            {"color_linenr",    "#315b31"},
            {"color_select",    "#7c9283"},
            {"color_type",      "#3e5f39"},
            {"color_accent",    "#0e1e0e"},
            {"color_string",    "#2b342a"},
            {"color_boolean",   "#293d29"},
            {"color_variable",  "#2b342a"},
            {"color_number",    "#445046"},
            {"color_comment",   "#5b5f59"},
            {"color_function",  "#0e1e0e"}
        }
    }
};

/** ======================================================================== */
int main() {
    int i, j;
    FILE *f_template;

    for (i = 0; i < N_EDITORS; i++) {
        f_template = fopen(editors[i].template, "r");
        if (f_template == NULL) {
            /* Unable to open template file, need to abort */
            printf("\n!!! Unable to open template file: '%s'\n",
                    editors[i].template);
            printf("Please check whether the file exists.\n\n");
            exit(EXIT_FAILURE);
        }

        /* Generate the editor themes using the template file
         * and a formatted output path */
        for (j = 0; j < N_THEMES; j++) {
            generate_theme(f_template, editors[i].output, themes[j]);
        }
    }

    fclose(f_template);
    return EXIT_SUCCESS;
}
