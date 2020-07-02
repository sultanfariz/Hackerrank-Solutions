#include <sstream>
#include <vector>
#include <iostream>
using namespace std;

vector<int> parseInts(string str) {
    vector<int> arr;
    stringstream ss(str);
    char ch=',';
    int n;
    while(ss>>n){
        arr.push_back(n);
        ss>>ch;
    }
    return arr;
}

int main() {
    string str;
    cin >> str;
    vector<int> integers = parseInts(str);
    for(int i = 0; i < integers.size(); i++) {
        cout << integers[i] << "\n";
    }
    
    return 0;
}
