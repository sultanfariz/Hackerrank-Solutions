import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
            Scanner sc=new Scanner(System.in);
            System.out.println("================================");
            for(int i=0;i<3;i++){
                String s1=sc.next();
                int x=sc.nextInt();
                //Complete this line
                int digit = 15-(s1.length());
                System.out.print(s1);
                for(int j=0; j<digit; j++) System.out.print(" ");
                if(x<100 && x>9) System.out.println("0"+x);
                else if(x<100 && x<10) System.out.println("00"+x);
                else System.out.println(x);
            }
            System.out.println("================================");

    }
}



