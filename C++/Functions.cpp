#include <iostream>
#include <cstdio>
#include <vector>
using namespace std;

/*
Add `int max_of_four(int a, int b, int c, int d)` here.

*/
int max_of_four(int a, int b, int c, int d){
    int m=0,n=0,x=0;
    if(a>b) m=a;
    else m=b;
    if(c>d) n=c; else n=d;
    if(m>n) return m;
    else return n;
}

int main() {
    int a, b, c, d;
    scanf("%d %d %d %d", &a, &b, &c, &d);
    int ans = max_of_four(a, b, c, d);
    printf("%d", ans);
    
    return 0;
}
