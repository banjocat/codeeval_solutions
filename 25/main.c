#include <stdio.h>


int main(int argc, char** argv)
{
   for (int i = 1; i <= 99; i++) {
      if (i == 1 || i % 2 != 0)
         printf("%d\n", i);
   }
}
