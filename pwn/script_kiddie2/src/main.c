#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// バッファリングを無効化して時間制限を60秒に設定
__attribute__((constructor)) void setup() {
  alarm(60);
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
}

void show_flag() {
  char flag_name[5] = {};
  printf("Which flag do you want?");
  read(0, flag_name, 5);

  char cmd[0x10];
  sprintf(cmd, "echo %s", flag_name);

  // show FLAG
  system(cmd);
}

int main() { show_flag(); }
