#include <iostream>
#include <sstream>
using namespace std;

/*
Enter code for class Student here.
Read statement for specification.
*/

class Student{
    private:    
        int a; string f; string l; int s;

    public:
        void set_age(int age){a=age;}
        void set_standard(int standard){s=standard;}
        void set_first_name(string first_name){f=first_name;}
        void set_last_name(string last_name){l=last_name;}
        int get_age(){return a;}
        int get_standard(){return s;}
        string get_first_name(){return f;}
        string get_last_name(){return l;}
        void to_string(){
            cout<<a<<",";
            cout<<f<<",";
            cout<<l<<",";
            cout<<s<<endl;
        }
};


int main() {
    int age, standard;
    string first_name, last_name;
    
    cin >> age >> first_name >> last_name >> standard;
    
    Student st;
    st.set_age(age);
    st.set_standard(standard);
    st.set_first_name(first_name);
    st.set_last_name(last_name);
    
    cout << st.get_age() << "\n";
    cout << st.get_last_name() << ", " << st.get_first_name() << "\n";
    cout << st.get_standard() << "\n";
    cout << "\n";
    // cout << 
    st.to_string();
    
    return 0;
}
