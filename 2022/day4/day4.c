#include <stdio.h>

int main() {
    int s1, s2, e1, e2;
    int err_code;
    int counter_in = 0;
    int counter_overl = 0;

    while (1) {
        err_code = scanf("%d-%d,%d-%d", &s1, &e1, &s2, &e2);
        if (err_code == -1) {
            break;
        }
        int first_in_second = s1 >= s2 && e1 <= e2;
        int second_in_first = s2 >= s1 && e2 <= e1;
        if (first_in_second || second_in_first) {
            counter_in++;
        } 
        int first_over_second = s1 <= e2 && e1 >= s2;
        int second_over_first = s2 <= e1 && e2 >= s1;
        if (first_over_second || second_over_first) {
            counter_overl++;
        }
    }
    printf("Contained sequences: %d\n", counter_in);
    printf("Overlapping sequences: %d\n", counter_overl);
}