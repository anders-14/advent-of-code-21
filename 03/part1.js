const fs = require('fs');

let input = [];

try {
  input = fs.readFileSync('./input', 'utf8').split('\n');
} catch (err) {
  console.error(err);
}

const bitCount = new Array(12).fill(0);

for (const binNum of input) {
  for (const idx in binNum) {
    bitCount[idx] += binNum[idx] == 0 ? 0 : 1; 
  }
}

const gamma = parseInt(bitCount.map(a => a / input.length < 0.5 ? 0 : 1).join(''), 2);
console.log("Gamma:             " + gamma);

const epsilon = parseInt(bitCount.map(a => a / input.length < 0.5 ? 1 : 0).join(''), 2);
console.log("Epsilon:           " + epsilon);

console.log("Power consumption: ", gamma * epsilon);
