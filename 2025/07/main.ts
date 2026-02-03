import fs from 'node:fs/promises';
// const fs = require('node:fs/promises');

async function read(path: string) {
  const data = await fs.readFile(path, { encoding: 'utf8' });
  return data
}

class Beam {
  y: number = -1
  x: number = -1
  counter: number = 1

  constructor(y: number, x: number, counter: number) {
    this.y = y
    this.x = x
    this.counter = counter
  }

  updateCounter(num: number) {
    this.counter = num
  }
  moveDown() {
    this.y++
  }
  moveLeft() {
    this.x--
  }
  moveRight() {
    this.x++
  }
}

function removeDups(beams: Beam[]) {
  for (let i = 0; i < beams.length - 1; i++) {
    for (let j = i + 1; j < beams.length; j++) {
      if (beams[i].y === beams[j].y && beams[i].x === beams[j].x) {
        beams[i].updateCounter(beams[i].counter + beams[j].counter)
        beams = beams.filter((_, index) => index !== j)
        j--
      }
    }
  }
  return beams
}

function solveAllParts(data: string[][]) {
  let startIndex = data[0].indexOf('S')
  let beams = [new Beam(0, startIndex, 1)]
  let split: { y: number, x: number }[] = []

  while (beams[0].y !== data.length - 1) {
    // move down
    let newBeams: Beam[] = []
    for (let b of beams) {
      b.moveDown()
      if (data[b.y][b.x] === '^') {
        split.push({ y: b.y, x: b.x, })
        let b2 = new Beam(b.y, b.x, b.counter)
        b.moveLeft()
        b2.moveRight()
        newBeams.push(b2)
      }
    }
    // combine
    beams = [...beams, ...newBeams]
    // remove dups
    beams = removeDups(beams)
  }

  let sum = 0
  for (let b of beams) {
    sum += b.counter
  }

  console.log('splits amount:', split.length)
  console.log('sum:', sum)
}

async function main() {
  // let data: string = await read('example.txt')
  let data: string = await read('input.txt')
  let arr = data.split('\n').map(row => row.split(''))
  solveAllParts(arr)
}

main()
