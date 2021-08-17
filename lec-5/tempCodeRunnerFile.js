// console.log("Bài 1")
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
// isPalindrome("eye");
// isPalindrome("_eye")
// isPalindrome("race car")
// isPalindrome("not a palindrome")
// isPalindrome("a man, a plan, a canal. Panama")
// isPalindrome("my age is 0,0 si ega ym")
// isPalindrome("0-0 (: /-\ :) 0-0")
// isPalindrome("five|\_/|four")

console.log("Bài 2")
function uniqueUnion(...args){
    var i =0;
    var s = new Array;
    s[0]= args[0][0];
    for(i;i<args.length;i++){
       for(var k = 0;k<args[i].length;k++){
        //console.log(args[i][k]);

        console.log(":"+args[0][0])
        for(var j=0;s.length;j++){
            console.log(s.length);

            if(s[j]!=args[i][k]){
                s[s.length]= args[i][k];
                //console.log("add");
            }
           
        }
        //console.log(s.length);
       }
        
    }
    console.log(s)
}
uniqueUnion([1, 3, 2], [5, 2, 1, 4], [2, 1])