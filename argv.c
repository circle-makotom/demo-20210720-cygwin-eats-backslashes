#include <stdio.h>

int main(int argc, char *argv[])
{
    int iter = 0;

    for (iter = 0; iter < argc; iter += 1)
    {
        printf("%s\n", argv[iter]);
    }
}
