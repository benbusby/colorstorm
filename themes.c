#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <themes.h>

/**
 * Function: generateTheme
 * --------------------
 * Generates a set of themes using a particular editor template 
 * (Vim, VSCode, etc).
 *
 * fTemplate: the editor theme template file
 * output: the char path to a theme
 *         (note: these are meant to be formatted with the name of the theme
 *         i.e. vim/colors/%s.vim -> vim/colors/earthbound.vim)
 * theme: the Theme object to use for creating the theme files
 */
void generateTheme(FILE *fTemplate, const char *output, struct Theme theme) {
    // File pointer to hold reference of input file
    FILE *fOutput;
    FILE *fOutputDark;

    char editorTheme[64] = {0}, editorDarkTheme[64] = {0};
    char name[32] = {0}, darkName[32] = {0};
    char darker[] = "-darker";

    sprintf(name, "%s", theme.name);
    sprintf(editorTheme, output, name);
    fOutput = fopen(editorTheme, "w");

    writeOutput(fTemplate, fOutput, theme);

    // Generate modified "darker" theme if eligible
    if (theme.darkMode) {
        char temp[strlen(theme.values[SWAP_IDX_A][1])];
        strcpy(temp, theme.values[SWAP_IDX_A][1]);
        theme.values[SWAP_IDX_B][1] = theme.values[SWAP_IDX_B][1];
        theme.values[SWAP_IDX_B][1] = temp;

        // Update name of output file
        strcat(darkName, name);
        strcat(darkName, darker);

        sprintf(editorDarkTheme, output, darkName);
        fOutputDark = fopen(editorDarkTheme, "w");

        writeOutput(fTemplate, fOutputDark, theme);
    }
}

/**
 * Function: writeOutput
 * --------------------
 * Writes the contents of a template theme into a completed editor syntax theme
 *
 * fInput: the editor theme template file
 * fOutput: the theme output file to write
 * theme: a Theme object to use for replacing generic fields with theme
 *        specific values
 */
void writeOutput(FILE *fInput, FILE *fOutput, struct Theme theme) {
    int i;
    char buffer[BUFFER_SIZE] = {0};

    if (fOutput == NULL) {
        /* Unable to open output file, need to abort */
        printf("\n!!! Unable to open output file\n");
        printf("Please check whether the path exists.\n\n");
        exit(EXIT_FAILURE);
    }

    // Iterate over file contents and replace all template values
    while ((fgets(buffer, BUFFER_SIZE, fInput)) != NULL) {
        for (i = 0; i < N_VALUES; i++) {
            // Replace occurrences of field with new value
            replaceAll(buffer, theme.values[i][0], theme.values[i][1]);
        }

        fputs(buffer, fOutput);
    }

    rewind(fInput);
}

/**
 * Function: replaceAll
 * --------------------
 * Replaces all instances of one word with another.
 *
 * str: the full string to read for matches
 * oldWord: the word in the string to replace
 * newWord: the new word to use in place of oldWord
 */
void replaceAll(char *str, const char *oldWord, const char *newWord) {
    char *pos, temp[BUFFER_SIZE];
    int index = 0;

    while ((pos = strstr(str, oldWord)) != NULL) {
        strcpy(temp, str);
        index = pos - str;

        // Cut string at old word pos, append new word
        str[index] = '\0';
        strcat(str, newWord);
        
        // Add remainder of the string back
        strcat(str, temp + index + strlen(oldWord));
    }
}

/** ======================================================================== */
int main() {
    int i, j;
    FILE *fTemplate;

    for (i = 0; i < N_EDITORS; i++) {
        fTemplate = fopen(editors[i].template, "r");
        if (fTemplate == NULL) {
            /* Unable to open template file, need to abort */
            printf("\n!!! Unable to open template file: '%s'\n",
                    editors[i].template);
            printf("Please check whether the file exists.\n\n");
            exit(EXIT_FAILURE);
        }

        /* Generate the editor themes using the template file
         * and a formatted output path */
        for (j = 0; j < N_THEMES; j++) {
            generateTheme(fTemplate, editors[i].output, themes[j]);
        }
    }

    fclose(fTemplate);
    return EXIT_SUCCESS;
}
