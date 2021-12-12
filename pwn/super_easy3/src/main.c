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
  double rate;
  time_t deadline;
} TASK;

TASK task;

void show_flag() {
  double diff = difftime(time(NULL), task.deadline);
  if (task.is_done && task.score == 0x1337 && 0 >= diff &&
      diff >= 60 * 60 * 24) {
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

  // 締め切りは次の日に設定
  time_t now = time(NULL);
  struct tm *now_tm = localtime(&now);
  now_tm->tm_mday += 1;
  task.deadline = mktime(now_tm);
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
  printf("task score: %d\n", task.score);
  printf("task deadline: %s\n", ctime(&task.deadline));

  double diff = difftime(time(NULL), task.deadline);
  if (task.is_done && task.score == 0x1337 && 0 < diff && diff < 60 * 60 * 24) {
    show_flag();
  }
}
