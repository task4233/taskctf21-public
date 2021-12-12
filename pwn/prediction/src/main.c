#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

// バッファリングを無効化して時間制限を60秒に設定
__attribute__((constructor)) void setup() {
  alarm(60);
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
}

char binsh[0x8] = "/bin/sh";

void debug_stack_dump(unsigned long rsp, unsigned long rbp) {
  unsigned long i;
  puts("\n***start stack dump***");
  i = rsp;
  while (i <= rbp + 32) {
    unsigned long *p;
    p = (unsigned long *)i;
    printf("0x%lx: 0x%016lx", i, *p);
    if (i == rsp) {
      printf(" <- rsp");
    } else if (i == rbp) {
      printf(" <- rbp");
    } else if (i == rbp + 8) {
      printf(" <- return address");
    }
    printf("\n");
    i += 8;
  }
  puts("***end stack dump***\n");
}

void show_flag() {
  char flag_name[0x30];
  printf("What is the flag?");
  flag_name[read(0, flag_name, 0x100) - 1] = '\0';

  // ここは正しいflagに差し替えられています
  if (strncmp(flag_name, "taskctf{", 8) != 0) {
    write(2, "Invalid flag", 13);
    return;
  }

  if (strcmp(flag_name, "taskctf{r0p_1s_f4mous_way}") == 0) {
    // show FLAG :)
    system(binsh);
  }

  // DEBUG MODE
  // TODO: remove
  {
    register unsigned long rsp asm("rsp");
    register unsigned long rbp asm("rbp");
    debug_stack_dump(rsp, rbp);
  }
}

int main() { show_flag(); }
