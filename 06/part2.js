const fs = require('fs');

let input = [];

try {
  input = fs.readFileSync('./input', 'utf8').split('\n')[0].split(',').map(a => +a);
} catch (err) {
  console.error(err);
}

const fishCount = new Array(9).fill(0);

for (let i of input) fishCount[i]++;

for (let i = 0; i < 256; i++) {
  fishCount.push(fishCount.shift());
  fishCount[6] += fishCount[8];
}

console.log("Lanternfish after 256 days:", fishCount.reduce((x, y) => x + y));

