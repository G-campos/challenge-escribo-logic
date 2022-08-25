const readline = require('node:readline');
const { stdin: input, stdout: output } = require('node:process');

const rl = readline.createInterface({ input, output });

const groupByThree = []
let sumGroupByThree = 0

const groupByFive = []
let sumGroupByFive = 0

let total = 0

const divisibleByThree = (input) => {
    return (input % 3) == 0 ? true : false
}

const divisibleByFive = (input) => {
    return (input % 5) == 0 ? true : false
}

rl.question('Informe um numero inteiro: ', (input) => {

  console.log(`\n`)

  let baseCalc = parseInt(input)

  for (let i = 0; i < baseCalc; i++) {
    if (divisibleByThree(i)) {
        groupByThree.push(i)
        sumGroupByThree += i
    }

    if (divisibleByFive(i)) {
        groupByFive.push(i)
        sumGroupByFive += i
    }
  }
  
  total = sumGroupByFive + sumGroupByThree
  
  console.log(`Número divisiveis por 3: ${groupByThree}`)
  console.log(`A soma dos numeros divisiveis por 3 é ${sumGroupByThree} \n`)

  console.log(`Número divisiveis por 5: ${groupByFive}`)
  console.log(`A soma dos numeros divisiveis por 5 é ${sumGroupByFive} \n`)

  console.log(`A soma Total é ${total} \n`)


  rl.close();
});

