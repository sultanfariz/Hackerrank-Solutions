import java.io.*;
import java.util.*;

public class Solution {

    public static void main(String[] args) {
        
        Scanner sc=new Scanner(System.in);
        String A=sc.next();
        /* Enter your code here. Print output to STDOUT. */
        int n=(A.length())-1; int x=1;
        for(int i=0; i<=n; i++){
            if(A.charAt(i)!=A.charAt(n)) x=0;
            n--;
        }
        if(x==0) System.out.println("No");
        else System.out.println("Yes");
    }
}



