#include <stdio.h>
#include <stdlib.h>

enum Colors {
    red, 
    green,
    purple,
    blue,
    white,
};

int main() {
    srand(time(NULL));

    int randomColors = rand() % 5;
}
