class Employee {
  int id;
  String name;
  float salary;
  void insertRecord(int i, String n, float s) {
    id = i;
    name = n;
    salary = s;
  }
  void display(){System.out.println(id+" "+name+" "+salary);}

}


public class TestEmployee {

  public static void main(String args[]){
     new Employee().display();  
    // e1.insertRecord(101,"ajeet",45000);  
    // e1.display();  
  }

}