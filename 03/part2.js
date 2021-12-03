const fs = require('fs');

let input = [];

try {
  input = fs.readFileSync('./input', 'utf8').split('\n').slice(0, -1);
} catch (err) {
  console.error(err);
}

const mostCommonBit = (arr, idx) => {
  const col = arr.map(a => a[idx]);

  const count0 = col.filter(a => a == '0').length;
  const count1 = col.filter(a => a == '1').length;

  return count0 <= count1 ? 1 : 0;
}

let oxygen_arr = [...input];
let co2_arr = [...input];

for (let i = 0; i < 12; i++) {
  if (oxygen_arr.length != 1) {
    const mostCommon = mostCommonBit(oxygen_arr, i);
    oxygen_arr = oxygen_arr.filter(a => a[i] == mostCommon);
  }
  if (co2_arr.length != 1) {
    const mostCommon = mostCommonBit(co2_arr, i);
    co2_arr = co2_arr.filter(a => a[i] != mostCommon);
  }
}

const o2 = parseInt(oxygen_arr[0], 2);
const co2 = parseInt(co2_arr[0], 2);

console.log("Oxygen:              " + o2);
console.log("CO2:                 " + co2);
console.log("Life support rating: " + o2 * co2);
