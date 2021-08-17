// console.log("Bài 2")
// function isPalindrome(n){
//     var i;
//     n = n.replace(/[\.,\?/:;{}[\]()!\s]/g,'');

//     for (i= 0 ;i< n.length/2;i++){
//         //console.log(n.length);
//         if(n[i]!=n[n.length-i-1]){
//             //console.log(n[i]+" "+n[n.length-i-1]);
//             console.log(n+": chuỗi không đối xứng");
//             return ;
//         }
//     } 
//     console.log(n+": chuỗi đối xứng");
// }
// //test case bài 2
// isPalindrome("eye");
// isPalindrome("_eye")
// isPalindrome("race car")
// isPalindrome("not a palindrome")
// isPalindrome("a man, a plan, a canal. Panama")
// isPalindrome("my age is 0,0 si ega ym")
// isPalindrome("0-0 (: /-\ :) 0-0")
// isPalindrome("five|\_/|four")

// console.log("Bài 3")
// function uniqueUnion(...args){
//     var i =0;
//     let s = [];
//     s.push(args[0][0]);
//     for(i;i<args.length;i++){
//        for(var k = 0;k<args[i].length;k++){
//         //console.log(args[i][k]);
//         //console.log("length"+s.length);
//             //console.log(s.length);
//         if(isempty(s,args[i][k])==0){
//             s.push(args[i][k]);
//         }
        
//         //console.log(s.length);
//        }
        
//     }
//     console.log(s)
// }
function isempty(s,x){
    
    for(var i=0;i<s.length;i++){
        if(s[i]==x){
            return 1;
        }
    }
    return 0;
}
// //test case bài 3
// uniqueUnion([1, 3, 2], [5, 2, 1, 4], [2, 1])
// uniqueUnion([1, 2, 3], [5, 2, 1])
// uniqueUnion([1, 2, 3], [5, 2, 1, 4], [2, 1], [6, 7, 8])

// console.log("Bài 4")
// function seekAndDestroy(x,...args){
//     //let s = []
//     for(var i = 0;i<args.length;i++){
//         for(var j =0;j<x.length;j++){
//             if(x[j]==args[i]){
//                 for(var k =j;k<x.length;k++){
//                     x[k]= x[k+1]
//                 }
//                 x.length--;
//             }
//         }
        
//     }
//     console.log(x)

// }
// //case bài 4
// seekAndDestroy([1, 2, 3, 1, 2, 3], 2, 3)
// seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3)
// seekAndDestroy(["foo","bar",1],"foo", 1)

console.log("Bài 5")
function toSpinalCase(s){
    for(var i =0;i<s.length;i++){
        if(s[i]==s[i].toUpperCase()){
            s.length++;
            for(var j=s.length;j>=i;j--){
                //var temp = s[j+1]
               
                if(i==j){
                    s[i]="-"
                }else{
                    s[j]=s[j-1]
                }
                
            }
            console.log(s[i])
        }
       // console.log(s[i])
    }
    // if(str[i]>=65&&str[i]<=90)
    //      str[i]=str[i]+32;
    console.log(s)
}
toSpinalCase("MyNameIsQuan")