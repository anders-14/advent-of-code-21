const fs = require('fs');

let input = [];

try {
  input = fs.readFileSync('./input', 'utf8').split('\n')[0].split(',').map(a => +a);
} catch (err) {
  console.error(err);
}

const simDay = (pop) => {
  const l = pop.length;
  for (let i = 0; i < l; i++) {
    if (pop[i] == 0) {
      pop[i] = 6;
      pop.push(8);
    } else {
      pop[i]--;
    }
  }
}

for (let i = 0; i < 80; i++) {
  simDay(input);
}

console.log('Lanternfish after 80 days:', input.length);
