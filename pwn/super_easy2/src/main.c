#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX_NAME_SIZE 16
#define MAX_TASK_SIZE 16

// バッファリングを無効化して時間制限を60秒に設定
__attribute__((constructor)) void setup() {
  alarm(60);
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
}

// TASKはタスクを管理する構造体
typedef struct {
  char name[MAX_NAME_SIZE]; // 最大16文字
  u_int32_t is_done;
  u_int32_t score;
} TASK;

TASK task;

void show_flag() {
  if (task.is_done == 0 || task.score != 0x1337) {
    printf("DO NOT CHEATING");
    return;
  }

  FILE *fp = fopen("flag", "r");
  if (fp == NULL) {
    fprintf(stderr, "flagが存在しません");
    return;
  }

  char flag[64];
  fgets(flag, 64, fp);
  printf(flag);

  fclose(fp);
}

void initialize() {
  // デフォルトではis_doneは0
  task.is_done = 0;

  // デフォルトのスコアは10
  task.score = 10;
}

int main() {
  // シグナルは無効化するよ
  struct sigaction sa = (struct sigaction){0};
  sa.sa_handler = SIG_IGN;
  if (sigaction(SIGINT, &sa, NULL) < 0) {
    return 1;
  }

  printf("Input task name\n");
  scanf("%s%*c", task.name);

  printf("task\n");
  printf("task name: %s\n", task.name);
  printf("task done: %d\n", task.is_done);

  if (task.is_done && task.score == 0x1337) {
    show_flag();
  }
}
