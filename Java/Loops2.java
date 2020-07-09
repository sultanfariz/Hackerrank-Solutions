import java.util.*;
import java.io.*;

class Solution{
    public static void main(String []argh){
        Scanner in = new Scanner(System.in);
        int t=in.nextInt();
        for(int i=0;i<t;i++){
            int a = in.nextInt();
            int b = in.nextInt();
            int n = in.nextInt();
            int ct=1;
            int prev=a+(ct*b);
            System.out.print(prev+" ");
            for(int j=2; j<=n; j++){
                ct*=2;
                prev+=((ct*b));
                System.out.print(prev+" ");
            }
            System.out.println();
        }
        in.close();
    }
}
