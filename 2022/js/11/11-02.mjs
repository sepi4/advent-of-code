import { readInput } from "../../lib/js/index.mjs"


class Monkey {
  inspected = 0
  index
  items = []
  divisibles = []

  constructor(m) {
    let indexMatch = m.match(/Monkey (\d+):\n/)
    let itemsMatch = m.match(/Starting items: (.+)\n/)
    let opMatch = m.match(/Operation: new = (.+)\n/)
    let divMatch = m.match(/Test: divisible by (\d+)/)
    let trueMatch = m.match(/If true: throw to monkey (\d+)/)
    let falseMatch = m.match(/If false: throw to monkey (\d+)/)

    this.index = indexMatch[1]
    this.items = itemsMatch[1].split(', ').map(x => Number(x))

    this.operation = function (old) {
      if (opMatch[1] === 'old * old') {
        return old * old
      } else if (opMatch[1] === 'old + old') {
        return old + old
      } else {
        let s = opMatch[1].substring(4, 5)
        let d = opMatch[1].substring(6)
        if (s === '+') {
          return old + Number(d)
        } else {
          return old * Number(d)
        }
      }
    }
    this.div = Number(divMatch[1])
    this.getNextIndex = function (num) {
      this.inspected++
      return num === 0 ? trueMatch[1] : falseMatch[1]
    }
  }
}

class D {
  xxx = {}
  constructor(num, monkeys) {
    for (let [k, m] of Object.entries(monkeys)) {
      this.xxx[m.index] = num % m.div
    }
  }
}

// managed to solve this second part only because my GF is math teacher :)
// she had idea for algorithm
async function main() {
  let input = await readInput('./2022/11/input.txt')
  // let input = getExampleInput()
  let mArr = input.split(/\n\n/)

  let monkeys = {}
  for (let m of mArr) {
    let monkey = new Monkey(m)
    monkeys[monkey.index] = monkey
  }

  for (let [k, m] of Object.entries(monkeys)) {
    for (let x of m.items) {
      let d = new D(x, monkeys)
      m.divisibles.push({
        startingValue: x,
        d,
      })
    }
  }

  let round = 10000
  // let round = 1000
  // let round = 100
  // let round = 20
  // let round = 1
  let n = 0
  while (++n <= round) {
    for (let [k, m] of Object.entries(monkeys)) {
      m.divisibles = m.divisibles.reverse()
      while (m.divisibles.length) {
        let obj = m.divisibles.pop()
        for (let [k, v] of Object.entries(obj.d.xxx)) {
          let newValue = m.operation(v) % monkeys[k].div
          obj.d.xxx[k] = newValue
        }
        let nextIndex = m.getNextIndex(obj.d.xxx[k])
        monkeys[nextIndex].divisibles.push(obj)
      }
    }
  }

  let ordered = Object.values(monkeys).sort((a, b) => b.inspected - a.inspected)
  console.log(ordered[0].inspected * ordered[1].inspected)
  // console.log(monkeys)
}

main()

function getExampleInput() {
  return `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
}
