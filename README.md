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

