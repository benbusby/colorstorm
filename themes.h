#include <string.h>
#define N_THEMES 7
#define N_VALUES 16
#define N_EDITORS 4
#define SWAP_IDX_A 2
#define SWAP_IDX_B 4
#define BUFFER_SIZE 1000

typedef struct Theme {
    char *name;
    int dark_mode;
    char *values[N_VALUES][2];
} Theme;

typedef struct Editor {
    char *template;
    char *output;
} Editor;

void replace_all(char *str, const char *old_word, const char *new_word);
void generate_theme(FILE *template, const char *output, struct Theme theme);
void write_output(FILE *f_input, FILE *f_output, struct Theme theme);

extern struct Editor editors[N_EDITORS];
extern struct Theme themes[N_THEMES];
