'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.trim().split('\n').map(string => {
        return string.trim();
    });
    
    main();    
});

function readLine() {
    return inputString[currentLine++];
}

function getLetter(s) {
    let letter;
    let a=["a","i","u","o","e"];
    let b=["b","c","d","f","g"];
    let c=["h","j","k","l","m"];
    // Write your code here

    switch(true){
        case a.includes(s[0]) :
            letter="A"; break;
        case b.includes(s[0]) :
            letter="B"; break;
        case c.includes(s[0]) :
            letter="C"; break;
        
        default: 
            letter="D";
    }
    
    return letter;
}

