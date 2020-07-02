#include <iostream>
#include <string>
#include <cstring>
using namespace std;

int main() {
	// Complete the program
    string as,bs;
    cin>>as>>bs;
    int an=as.length(), bn=bs.length();
    cout<<an<<" "<<bn<<endl;
    cout<<as<<bs<<endl;
    char temp=as[0];
    as[0]=bs[0]; bs[0]=temp;
    cout<<as<<" "<<bs;
    return 0;
}
