import fs from 'node:fs/promises';

async function read(path) {
  const data = await fs.readFile(path, { encoding: 'utf8' });
  return data
}

class Beam {
  y = null
  x = null

  constructor(y, x) {
    this.y = y
    this.x = x
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

function removeDups(beams) {
  for (let i = 0; i < beams.length - 1; i++) {
    for (let j = i + 1; j < beams.length; j++) {
      if (beams[i].y === beams[j].y && beams[i].x === beams[j].x) {
        beams = beams.filter((_, index) => index !== j)
      }
    }
  }
  return beams
}

function part1(data) {
  let startIndex = data[0].indexOf('S')
  let beams = [new Beam(0, startIndex)]
  let split = []

  while (beams[0].y !== data.length - 1) {
    // move down
    let newBeams = []
    for (let b of beams) {
      b.moveDown()
      if (data[b.y][b.x] === '^') {
        split.push({ y: b.y, x: b.x, })
        let b2 = new Beam(b.y, b.x)
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
  console.log(split.length)
}

function part2(matrix) {
}

async function main() {
  let data = await read('./07/input.txt')
  let arr = data.split('\n').map(row => row.split(''))
  part1(arr)
}

main()
