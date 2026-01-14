/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2026-1-7
 * @fileoverview
 */

let questions: string[] = [];
let answers: string[] = [];
let correct: number = 0;
let index: number = 0;

console.log("Welcome to Trivia!\nThis is a 2 player game where you make a mini quiz for your friend and they answer.\n Be sure to click enter after you find typing every input");
  let numQ=Number(
    prompt("How many questions would you like to quiz your friend on? Please input the number as a integer not as a word")
  );

  while (isNaN(numQ) || numQ <= 0) {
    numQ = Number(prompt("Please enter a valid number greater than 0."));
  }

let questionNum=1;
while (index < numQ) {
  let question = prompt("Enter question " + questionNum + ":");
  let answer = prompt("Enter the correct answer for the question:");
  if (question == null || answer == null || question == "" || answer == "") {
    console.log("Question and answer cannot be empty.");
  } else {
    questions[index] = question;
    answers[index] = answer;
    index = index + 1;
  }
}

let playIndex: number = 0;
console.log(" \n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
while (playIndex < numQ) {
  let userAnswer=prompt(questions[playIndex] + "?");

  if (userAnswer == answers[playIndex]) {
    console.log("Correct!");
    correct = correct + 1;
  } else {
    console.log(
      "Incorrect. The correct answer was: " + answers[playIndex]
    );
  }

  playIndex = playIndex + 1;
}

//Printing final results
console.log("Quiz complete. You got a score of " + correct + " out of " + numQ + ".");
