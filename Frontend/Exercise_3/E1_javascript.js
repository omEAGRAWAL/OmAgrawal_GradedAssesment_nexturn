// Basic Exercises

console.log("===== Basic Exercises =====");

// Exercise 1: Basic Arithmetic with Variables
let num1 = 15;
let num2 = 5;

console.log("Addition:", num1 + num2); // 20
console.log("Subtraction:", num1 - num2); // 10
console.log("Multiplication:", num1 * num2); // 75
console.log("Division:", num1 / num2); // 3

// Exercise 2: Data Type Check
let stringVar = "Hello, World!";
let numberVar = 42;
let booleanVar = true;
let objectVar = { name: "Om" };
let undefinedVar;

console.log("Type of stringVar:", typeof stringVar); // string
console.log("Type of numberVar:", typeof numberVar); // number
console.log("Type of booleanVar:", typeof booleanVar); // boolean
console.log("Type of objectVar:", typeof objectVar); // object
console.log("Type of undefinedVar:", typeof undefinedVar); // undefined

// Exercise 3: Conditionals
let number = -7;
if (number > 0) {
  console.log(`${number} is positive.`);
} else if (number < 0) {
  console.log(`${number} is negative.`);
} else {
  console.log(`${number} is zero.`);
}

// Exercise 4: Even or Odd using Loops
for (let i = 1; i <= 10; i++) {
  if (i % 2 === 0) {
    console.log(`${i} is even.`);
  } else {
    console.log(`${i} is odd.`);
  }
}

// Exercise 5: Sum of Numbers using a While Loop
let sum = 0;
let i = 1;
while (i <= 100) {
  sum += i;
  i++;
}
console.log("Sum of numbers from 1 to 100:", sum); // 5050

// Advanced Exercises

console.log("\n===== Advanced Exercises =====");

// Exercise 6: Advanced Arithmetic with Functions
function add(a, b) {
  return a + b;
}
function subtract(a, b) {
  return a - b;
}
function multiply(a, b) {
  return a * b;
}
function divide(a, b) {
  if (b === 0) return "Error: Division by zero";
  return a / b;
}
console.log("Addition:", add(8, 2)); // 10
console.log("Subtraction:", subtract(8, 2)); // 6
console.log("Multiplication:", multiply(8, 2)); // 16
console.log("Division:", divide(8, 2)); // 4

// Exercise 7: Data Types and Array Manipulation
let arr = [1, 2, 3];
arr.push(4);
console.log("After push:", arr); // [1, 2, 3, 4]
arr.pop();
console.log("After pop:", arr); // [1, 2, 3]
arr.shift();
console.log("After shift:", arr); // [2, 3]
arr.unshift(0);
console.log("After unshift:", arr); // [0, 2, 3]
arr.forEach((value) => console.log("Value:", value)); // 0, 2, 3

let multiDimArr = [
  [1, 2],
  [3, 4],
];
console.log("Multi-dimensional Array:", multiDimArr); // [[1, 2], [3, 4]]

// Exercise 8: Complex Conditional Logic
let age = 25;
let experience = 3; // in years
let skills = ["JavaScript", "React", "Node.js"];

if (age >= 18 && experience >= 2 && skills.includes("React")) {
  console.log("Eligible for the job.");
} else {
  console.log("Not eligible for the job.");
}

// Exercise 9: Nested Loops and Prime Number Check
for (let num = 2; num <= 20; num++) {
  let isPrime = true;
  for (let j = 2; j < num; j++) {
    if (num % j === 0) {
      isPrime = false;
      break;
    }
  }
  if (isPrime) console.log(`${num} is a prime number.`);
}

// Exercise 10: FizzBuzz with Loops and Conditionals
for (let num = 1; num <= 100; num++) {
  if (num % 3 === 0 && num % 5 === 0) {
    console.log("FizzBuzz");
  } else if (num % 3 === 0) {
    console.log("Fizz");
  } else if (num % 5 === 0) {
    console.log("Buzz");
  } else {
    console.log(num);
  }
}
