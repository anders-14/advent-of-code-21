#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#if 0
#define FILENAME "sample"
#define INPUT_COUNT 10
#define GRID_SIZE 10
#else
#define FILENAME "input"
#define INPUT_COUNT 500
#define GRID_SIZE 1000
#endif

int grid [GRID_SIZE][GRID_SIZE] = {0};

typedef struct {
  int x1, y1, x2, y2;
} Line;

Line lines[INPUT_COUNT];
size_t lines_sz = 0;

int min(int a, int b) { return a < b ? a : b; }
int max(int a, int b) { return a < b ? b : a; }

int main() {
  FILE *fp = fopen(FILENAME, "r");

  size_t nread;
  char c;

  // There must be a better way to do this
  char num[3] = {0};
  size_t num_sz = 0;
  int start = 1;
  int x1, y1, x2, y2 = 0;
  while ((nread = fread(&c, 1, 1, fp) != 0) && lines_sz != INPUT_COUNT) {
    switch (c) {
      case ',':
        if (start == 1) {
          start = 0;
          x1 = atoi(num);
          memset(num, 0, 3);
          num_sz = 0;
        } else {
          start = 1;
          x2 = atoi(num);
          memset(num, 0, 3);
          num_sz = 0;
        }
        break;
      case ' ':
      case '-':
        break;
      case '>':
        y1 = atoi(num);
        memset(num, 0, 3);
          num_sz = 0;
        break;
      case '\n':
        y2 = atoi(num);
        memset(num, 0, 3);
        num_sz = 0;
        lines[lines_sz].x1 = x1;
        lines[lines_sz].y1 = y1;
        lines[lines_sz].x2 = x2;
        lines[lines_sz].y2 = y2;
        lines_sz++;
        x1 = 0;
        y1 = 0;
        x2 = 0;
        y2 = 0;
        break;
      default:
        num[num_sz++] = c;
    }
  }
  fclose(fp);

  for (int i = 0; i < lines_sz; i++) {
    Line *l = &lines[i];

    if (l->x1 == l->x2) {
      for (int j = 0; j <= max(l->y1, l->y2) - min(l->y1, l->y2); j++) {
        grid[j + min(l->y1, l->y2)][l->x1]++;
      }
    } else if (l->y1 == l->y2) {
      for (int j = 0; j <= max(l->x1, l->x2) - min(l->x1, l->x2); j++) {
        grid[l->y1][j + min(l->x1, l->x2)]++;
      }
    } else {
      int minx = min(l->x1, l->x2); 
      int miny = min(l->y1, l->y2); 
      int delta = max(l->x1, l->x2) - min(l->x1, l->x2);

      for (int j = 0; j <= delta; j++) {
        if (minx == l->x1) {
          if (miny == l->y1) {
            grid[l->y1 + j][minx + j]++;
          } else {
            grid[l->y1 - j][minx + j]++;
          }
        } else {
          if (miny == l->y2) {
            grid[l->y2 + j][minx + j]++;
          } else {
            grid[l->y2 - j][minx + j]++;
          }
        }
      }
    }
  }

  int count = 0;
  for (int i = 0; i < GRID_SIZE; i++) {
    for (int j = 0; j < GRID_SIZE; j++) {
      if (grid[i][j] >= 2) count++;
    }
  }

  printf("Count: %d\n", count);

  return 0;
}
