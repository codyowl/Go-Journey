# Go-Journey
My journal with classified tracks to master this beautiful programming language "GOlang"

## Variables

     <variable_name> <data_type> 

     number int

   
   - Using 'var':


    var <variable_name> <data_type>

    var number int

    var <variable_name_1>, <variable_name_2>, <variable_name_3> <common_data_type>

    var a,b,c string 
    
    
   - Using Short declaration:

    <my_variable> := <value>

    Name := "John Carmack"  

  
   - Type Inference:

     
    var <my_variable> = <value>
    
    var name = "straight away value"
     
   
   - Reassigning varaible: 

    var name string = "Linus Torvalds"

    name = "Steve Wozniak"

  - Multiple Variables:

        <varialbe_name1>, <variable_name2>, <variable_name3> <common_data_type>
        x, y, z int
        
        <variable_1>, <variable_2> := <value_1>, <value_2>
        myNumber, myName := 1, "some name"

        var <variable_1>, <variable_2> = <value_1>, <value_2>
        var myvarNumber, myvarName = 1, "var name"
  
  - Multiline variables:
  
        var (
                <variable_1> = <value_1>
                <variable_2> = <value_2>
            )
            
        var (
                 myName = "some name"
                 myNumber = 2
                 myBoolean = false
            )     
## Pointers:
   - Declaring pointer:
   
	var firstname *string 
	Here, firstname is a pointer with a value string

   - Dereferencing :
   
   	 
	*firstname = "john"
	
	var firstname *string = new(string)
	*firstname = "john"
	fmt.Println(*firstname)
            Note : we have to use dereferencing operator to denote a pointer in pringln we do have *firstname
	    
	
	firstname := "john"
	pointer := &firstname
	fmt.Println(*pointer)

	Now, changing firstname will reflect while printing the pointer variable
	firstname := "hai"
	fmt.Println(*pointer) => this will pring “hai”


	

## conditional statement:
   -if:
         
          if <condition> {
                    <do something>
                }
   -else if
          
          if <condition_for_if> {
                    <do something>
                }
          else if <condition_for_else_if> {
                    <do something>
                }
          else {
                    <do something when none of the above conditions met>
               }
               
    
   -switch :
          
           switch <condition> {
           case <condition_1>:
               <do_something_if_this_case_matches>
           case <condition_2>:
               <do_something_if_This_case_matches>
           default:
               <do_something_when_none_of_thecase_mathces>
           }
                  
## for loop use cases:
    
 - Using for like while loop:
    
       counter := 0
       for counter < 10 {
          fmt.Println(counter)  
       }

	
 - The typical for loop :
    
        for i := 0; i <= 10; i++ {
            fmt.Println(i)
        }

 - Infinite loop (like While True in python)
	

        for {
           fmt.Println("I will be printed infinitely ! ")
          
        }

## Data Structures:
### Array(Immutable):

   - Declaration :
     using var:
     
	      var <array_name> [<size>]<datatype>
	      var myNumberArray [6] int
	      
      using shorhand declartion:
      
              <array_name> := [<size>]<datatype>{value1, value2}  
              myColorArray := [4]string{"Red", "Green", "Blue", "Yellow"}
	      
	      
   - Accessing and assigning values by index:
             
	     //Assigning
             var programmer [3]string
             programmer[0] = "John Carmack"
             programmer[1] = "Sean Parker"
             programmer[2] = "Torvalds"
             
             //Accessing by index
	     fmt.Println(programmer[0]) 
	      
### slice(Mutable)

   - Declaration : 
     using var(Note : similar to array but we don't need to specify the size):
               
	       var <sice_name> []<data_type>
               var musicians []string
               func main() {
                     musicians = []string{"Bob dylan", "Jimi Hendrix", "Bob Marley"}
               }
	       
     using shorthand notation:
     
               <slice_name> := []<data_type>{value1, value2, value3}
               var musicians := []string{"Bob dylan", "Jimi Hendrix", "Bob Marley"} 
	       
     using make builtin functionn:
     
     	       <slice_name> := make([]<data_type>, <size>)
	       my_make_slice := make([]string, 3)
	       my_make_slice[0] = "first make slice value"
               my_make_slice[1] = "second make slice value"
	       my_make_slice[2] = "third make slice value"
   
### Map(Mutable)

   - Declaration :
     using var :
     
              var <map_name> = map[<key_data_type]<value_data_type>
              <map_name> = map[<key_data_type>]<value_data_type> {
                      key : value,
                      key : value,
                      key : value
               }
         
              var couples = map[string]string
	      func main() {
	           couples = map[string]string {
                      "Thala" : "Shalini",
                      "Thalaivar" : "Latha",
                      "Surya" : "Jyothika"
                          }
               }
 
       using shorthand declaration
       
               <map_name> := map[<key_data_type>]<value_data_type> {
                       key : value,
                       key : value,
                       key : value
                }

                founders := map[string]string{
                       "Drobbox" : "Drew",
                       "Valve" : " Gabe",
                       "Napster" : "Sean parker"
                }

   - Iterating a map(using range function)
     
                for key, value := range <map_name> {
                         fmt.Printf("The value of key is %v & value is %v", key, value)
                }
   
### Struct

   - Declaration :

                 type <struct_name> struct {
                         <varialbe_name_1> <datatype>
                         <variable_name_2> <datatype>
                  }

                 type Book struct {
                          authorName string
                          bookName string
			  genre string
                          numberOfPages int
                  }



   - Assigning :


                 <struct_variable> := <struct_name>{
                          <struct_field_1> : <value_1>
                          <struct_field_2> : <value_2>
                   }
		   
		 b := Book { 
                           authorName : "Josephine Cox",
                           bookName : "The Loner",
                           genre : "Domestic fiction",
                           numberOfPages : 345
                   }
		   

   - Accessing :
   
                 <struct_variable_name>.<struct_field_name>

                 fmt.Println(b.bookName)
		   
                             
## Functions

   
   - Declaration :
                 
	         func <function_name>(<parameter_name> <datatype>) <return_data_type> {
                                 <function_body>
                 }

   - Calling:
    
                 var <function_call_variable> <datatype>             


   Note : A function can return a list of values all the datatypes should be mentioned on return type.
   We don't need to specify function return type if its not returning any thing.

   Example of a simple addition function:

		 import "fmt"

		 func main(){
			var a int = 10
			var b int = 20
			var add_function_variable int

			add_function_variable = add(a, b)

			fmt.Printf("Sum value is : %d\n", add_function_variable)


		 }

		 # declaring function with parameter their datatype and return data type 	
			 func add(first_number, second_number, int) int {
				var result int
				result = first_number + second_number
				return result
			 }                


   - Multiple returns:

    
        package main

        import (
          "fmt"
        )

        func multipleReturn(x,y int)(int, int){
          return x + 2, y + 3
          }

        func main() {
          var a int = 10
          var b int = 20
          var funcVariable1, funcVariable2 = multipleReturn(a,b)
          fmt.Printf("Addition result of first variable: %d", funcVariable1)
          fmt.Printf("\nAddition result of second variable: %d", funcVariable2)
        }