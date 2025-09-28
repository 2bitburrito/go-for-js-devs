
export interface foo {
  numA: number;
  numB: number;
}

function sum(foos: foo[]) {
  let sum = 0
  for (let foo of foos) {
    sum += foo.numA
    sum += foo.numB
  }
}
