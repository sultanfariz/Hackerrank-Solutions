#include <iostream>

using namespace std;
/*
 * Create classes Rectangle and RectangleArea
 */

class Rectangle{
    protected:
        int width, height;
    public:
        void display(){
            cout<<width<<" "<<height<<endl;
        }
};

class RectangleArea:public Rectangle{
    int area;
    public:
        void display(){
            area=height*width;
            cout<<area<<endl;
        }
        void read_input(){
            // int a,b;
            // cin>>a>>b;
            // height=b; width=a;
            cin>>width>>height;
        }
};

int main()
{
    /*
     * Declare a RectangleArea object
     */
    RectangleArea r_area;
    
    /*
     * Read the width and height
     */
    r_area.read_input();
    
    /*
     * Print the width and height
     */
    r_area.Rectangle::display();
    
    /*
     * Print the area
     */
    r_area.display();
    
    return 0;
}
