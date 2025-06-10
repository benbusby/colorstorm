package main

const sampleText = `
 === Python ===
 {keyword}def{/keyword} {function}return_pi_or_nan{/function}(return_pi: {type}bool{/type}) -> {type}float{/type}:
     {function}print{/function}({string}"Hello world!"{/string})
     {comment}# This is a comment{/comment}
     pi: {type}float{/type} = {number}3.14159{/number}
     {keyword}if{/keyword} return_pi:
         {keyword}return{/keyword} pi
     {keyword}return{/keyword} {constant}numpy{/constant}.nan

 === Go ===
 {keyword}func{/keyword} {function}returnPiOrNaN{/function}(returnPi {type}bool{/type}) {type}float64{/type} {
         {constant}fmt{/constant}.{function}Println{/function}({string}"Hello world!"{/string})
         {comment}// This is a comment{/comment}
         pi := {number}3.14159{/number}
         {keyword}if{/keyword} returnPi {
                 {keyword}return{/keyword} pi
         }

         {keyword}return{/keyword} {constant}math{/constant}.{function}NaN(){/function}    
 }
`

func init() {
	// Validate the
}
