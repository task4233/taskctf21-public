#include <signal.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_NAME_SIZE 16

// バッファリングを無効化して時間制限を60秒に設定
__attribute__((constructor)) void setup() {
  alarm(60);
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
}

// TASKはタスクを管理する構造体
typedef struct {
  char name[MAX_NAME_SIZE];
  u_int32_t is_done;
  u_int32_t score;
} TASK;

TASK task;

void show_flag() {
  if (task.is_done == 0) {
    printf("DO NOT CHEATING");
    return;
  }

  FILE *fp = fopen("flag", "r");
  if (fp == NULL) {
    fprintf(stderr, "flag not found");
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
}

int main() {
  printf("Input task name\n");
  scanf("%s%*c", task.name);

  printf("task\n");
  printf("task name: %s\n", task.name);
  printf("task done: %d\n", task.is_done);

  // タスクの完了判定
  if (task.is_done) {
    show_flag();
  }
}
