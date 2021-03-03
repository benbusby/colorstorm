#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "themes.h"

/**
 * Function: generate_theme
 * --------------------
 * Generates a set of themes using a particular editor template 
 * (Vim, VSCode, etc).
 *
 * f_template: the editor theme template file
 * output: the char path to a theme
 *         (note: these are meant to be formatted with the name of the theme
 *         i.e. vim/colors/%s.vim -> vim/colors/earthbound.vim)
 * theme: the Theme object to use for creating the theme files
 */
void generate_theme(FILE *f_template, const char *output, struct Theme theme) {
    // File pointer to hold reference of input file
    FILE *f_output;
    FILE *f_output_dark;

    char editor_theme[64] = {0}, editor_dark_theme[64] = {0};
    char name[32] = {0}, dark_name[32] = {0};
    char darker[] = "-darker";

    sprintf(name, "%s", theme.name);
    sprintf(editor_theme, output, name);
    f_output = fopen(editor_theme, "w");

    write_output(f_template, f_output, theme);

    // Generate modified "darker" theme if eligible
    if (theme.dark_mode) {
        char temp[strlen(theme.values[SWAP_IDX_A][1])];
        strcpy(temp, theme.values[SWAP_IDX_A][1]);
        theme.values[SWAP_IDX_B][1] = theme.values[SWAP_IDX_B][1];
        theme.values[SWAP_IDX_B][1] = temp;

        // Update name of output file
        strcat(dark_name, name);
        strcat(dark_name, darker);

        sprintf(editor_dark_theme, output, dark_name);
        f_output_dark = fopen(editor_dark_theme, "w");

        write_output(f_template, f_output_dark, theme);
    }
}

/**
 * Function: write_output
 * --------------------
 * Writes the contents of a template theme into a completed editor syntax theme
 *
 * f_input: the editor theme template file
 * f_output: the theme output file to write
 * theme: a Theme object to use for replacing generic fields with theme
 *        specific values
 */
void write_output(FILE *f_input, FILE *f_output, struct Theme theme) {
    int i;
    char buffer[BUFFER_SIZE] = {0};

    if (f_output == NULL) {
        /* Unable to open output file, need to abort */
        printf("\n!!! Unable to open output file\n");
        printf("Please check whether the path exists.\n\n");
        exit(EXIT_FAILURE);
    }

    // Iterate over file contents and replace all template values
    while ((fgets(buffer, BUFFER_SIZE, f_input)) != NULL) {
        for (i = 0; i < N_VALUES; i++) {
            // Replace occurrences of field with new value
            replace_all(buffer, theme.values[i][0], theme.values[i][1]);
        }

        fputs(buffer, f_output);
    }

    rewind(f_input);
}

/**
 * Function: replace_all
 * --------------------
 * Replaces all instances of one word with another.
 *
 * str: the full string to read for matches
 * old_word: the word in the string to replace
 * new_word: the new word to use in place of old_word
 */
void replace_all(char *str, const char *old_word, const char *new_word) {
    char *pos, temp[BUFFER_SIZE];
    int index = 0;

    while ((pos = strstr(str, old_word)) != NULL) {
        strcpy(temp, str);
        index = pos - str;

        // Cut string at old word pos, append new word
        str[index] = '\0';
        strcat(str, new_word);
        
        // Add remainder of the string back
        strcat(str, temp + index + strlen(old_word));
    }
}

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
