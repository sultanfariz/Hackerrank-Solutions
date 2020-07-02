#include <stdio.h>

void update(int *a,int *b) {
    // Complete this function  
    int m=0,n=0,i=0,j=0;
    if((*a)>(*b)) {
        m=(*a); n=(*b);
    }else{
        m=(*b); n=(*a);
    }
    (*a)+=(*b);
    (*b)=m-n;
}

int main() {
    int a, b;
    int *pa = &a, *pb = &b;
    
    scanf("%d %d", &a, &b);
    update(pa, pb);
    printf("%d\n%d", a, b);

    return 0;
}
