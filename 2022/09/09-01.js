let input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
input = getInput()

let moves = input.split('\n').map(row => {
  row = row.split(' ')
  return {
    dir: row[0],
    n: Number(row[1]),
  }
})
// console.log(moves)

let H = { y: 0, x: 0, }
let T = { y: 0, x: 0, }

let tailMoves = {}

// console.log(H)
for (let move of moves) {
  while(move.n--) {
    H = makeHeadMove(H, move)
    T = makeTailMove(H, T)
    tailMoves[`${T.y},${T.x}`] = true
    // console.log(H, T)
  }
}

let tailMovesAmount = Object.keys(tailMoves).length
console.log(tailMovesAmount)

function makeTailMove(H, T) {
  let yDiff = Math.abs(H.y - T.y)
  let xDiff = Math.abs(H.x - T.x)
  if (yDiff > 1) {
    if (H.y > T.y) {
      T.y = T.y + 1
    } else {
      T.y = T.y - 1
    }
    if (xDiff > 0) {
      T.x = H.x
    }
  }
  if (xDiff > 1) {
    if (H.x > T.x) {
      T.x = T.x + 1
    } else {
      T.x = T.x - 1
    }
    if (yDiff > 0) {
      T.y = H.y
    }
  }
  return T
}

function makeHeadMove(coor, move) {
  switch (move.dir) {
    case 'U':
      return {
        y: coor.y + 1,
        x: coor.x,
      }
    case 'D':
      return {
        y: coor.y - 1,
        x: coor.x,
      }
    case 'L':
      return {
        y: coor.y,
        x: coor.x - 1,
      }
    case 'R':
      return {
        y: coor.y,
        x: coor.x + 1,
      }
    default:
      throw Error('wrong move')
  }
}

function getInput() {
  return `D 2
R 2
D 1
U 1
L 2
U 2
D 1
R 2
U 2
R 2
L 1
D 2
L 1
U 1
R 1
D 1
U 1
R 1
U 1
R 1
D 2
L 1
D 1
L 1
D 1
U 1
L 1
R 1
L 2
D 1
L 2
R 1
U 2
L 2
D 1
U 2
D 1
L 1
U 2
L 2
U 2
D 2
U 2
D 2
L 1
D 2
L 2
D 2
U 2
L 2
R 2
U 1
L 1
D 2
R 1
D 2
L 2
R 1
D 2
R 1
U 2
L 2
U 2
D 1
U 2
L 2
U 2
L 1
U 2
L 1
U 2
D 1
U 1
D 2
R 2
L 1
U 2
L 1
D 1
R 2
U 1
L 1
U 1
R 2
U 2
D 2
L 2
U 2
L 2
R 1
L 2
U 1
R 1
L 1
R 2
U 1
R 1
D 1
L 2
R 2
L 2
R 1
D 2
R 1
U 2
L 2
D 1
R 2
D 2
L 2
U 2
D 2
L 1
D 2
R 1
L 3
D 2
R 1
L 3
R 2
D 3
U 3
R 2
L 1
R 2
U 2
R 2
L 1
U 2
D 1
L 3
U 3
R 3
D 2
U 1
L 2
R 2
D 1
L 2
U 3
R 1
L 3
R 1
L 3
U 3
D 3
U 2
R 2
L 3
U 2
L 2
U 2
R 2
D 1
R 3
L 2
R 3
D 2
R 2
L 2
R 3
L 2
R 1
U 2
D 1
R 1
D 2
U 2
L 2
D 1
R 1
L 3
U 1
R 3
L 2
D 1
U 1
L 2
U 1
D 2
R 1
U 2
D 3
R 3
L 2
R 2
U 2
L 3
U 3
L 3
R 2
U 1
D 1
U 2
R 3
D 3
L 1
D 2
U 1
D 2
R 3
U 2
D 1
L 1
D 3
R 1
L 1
U 2
L 3
D 1
U 2
R 1
D 1
R 3
D 1
R 1
L 1
R 2
L 1
D 1
R 2
U 1
L 4
U 2
D 2
L 4
R 1
L 3
D 3
R 3
L 4
D 1
R 1
U 4
L 1
D 3
R 4
U 3
L 3
U 4
D 3
L 2
D 2
R 1
D 3
U 1
R 4
U 2
D 1
L 4
U 4
L 1
U 1
L 3
D 2
L 2
R 3
U 3
D 3
U 2
R 3
D 2
L 1
U 2
D 3
R 1
D 2
U 4
R 4
U 3
R 3
L 1
U 4
R 4
L 1
R 4
L 4
R 1
L 1
R 3
U 2
R 3
L 3
D 2
R 1
D 2
U 2
R 4
U 2
R 4
L 1
R 3
U 4
L 3
R 2
U 1
R 2
D 4
L 1
U 4
R 2
L 3
U 2
L 3
U 1
D 1
L 4
U 2
L 2
U 4
D 2
U 3
R 4
L 1
D 1
R 2
D 4
U 3
R 1
L 2
R 4
D 4
L 3
D 1
R 2
L 3
U 1
L 2
D 3
L 2
R 1
L 2
R 2
L 4
U 1
L 5
R 2
D 5
L 4
R 2
D 3
L 5
U 4
R 5
L 1
D 4
R 2
D 4
U 3
R 1
D 1
U 5
R 2
D 2
R 5
D 3
L 1
D 4
R 2
L 5
U 4
R 1
D 1
U 1
R 1
D 3
U 1
L 4
U 2
D 2
R 1
U 2
L 3
R 5
D 1
L 4
R 4
D 5
L 5
D 5
U 1
L 2
D 3
L 3
D 1
U 2
D 3
U 1
L 3
R 1
D 2
U 4
R 4
D 2
U 1
L 1
R 1
U 1
R 4
D 4
L 2
U 4
D 1
R 4
L 2
U 3
D 4
R 2
U 1
R 2
U 5
D 2
U 4
L 2
R 2
L 1
U 2
R 5
U 5
R 3
L 2
R 4
L 3
R 3
L 2
U 4
L 2
U 2
D 3
U 3
L 3
R 2
U 2
D 5
U 4
R 3
L 5
R 1
U 4
L 4
U 4
D 1
U 2
D 4
R 4
U 5
R 2
L 2
R 2
L 3
R 1
D 5
R 5
L 4
U 2
D 6
U 3
R 6
L 1
D 3
U 2
R 1
D 3
L 6
R 5
L 6
U 5
R 4
D 1
U 2
R 3
U 1
R 6
D 1
L 4
R 2
L 1
U 6
L 1
U 3
R 3
U 4
D 5
R 2
U 2
D 3
R 1
U 5
D 2
L 3
U 3
D 2
U 4
D 4
R 3
L 3
R 3
D 1
R 2
L 4
D 2
L 5
U 5
D 2
L 5
U 1
L 2
U 2
L 5
R 1
L 2
U 2
R 3
U 5
R 1
L 2
R 3
D 3
R 3
L 6
D 1
L 2
R 1
U 4
D 3
R 6
L 4
U 2
L 3
D 4
L 3
D 5
L 3
D 4
U 3
D 4
U 2
D 6
U 4
R 2
L 2
R 1
L 4
R 6
D 5
L 2
D 6
L 4
R 4
U 6
L 6
D 6
R 4
L 2
R 2
U 5
L 6
D 7
U 6
R 1
D 4
R 2
D 7
L 4
U 7
D 4
L 6
D 7
U 6
R 1
U 1
R 4
L 7
U 1
D 2
U 2
R 3
L 6
R 7
L 5
D 5
U 2
D 1
U 4
R 6
D 5
R 6
U 5
R 7
U 7
R 4
L 1
D 1
U 3
D 1
L 5
U 4
R 2
L 3
D 3
U 2
R 2
D 1
R 6
U 5
D 5
R 7
U 2
L 4
D 3
R 2
L 1
D 1
R 2
D 7
U 6
D 1
R 2
D 7
U 5
R 6
D 1
L 5
U 3
L 6
D 3
L 5
D 2
L 5
R 1
D 7
L 3
U 3
R 6
D 4
R 2
U 1
L 2
R 1
L 3
D 3
L 6
U 7
R 4
L 4
U 5
L 7
D 1
R 4
U 6
D 3
R 2
L 1
D 4
R 5
D 3
L 6
R 6
U 4
R 7
D 2
U 3
L 2
D 3
L 5
D 1
U 5
D 8
L 8
U 2
L 8
U 2
D 3
U 1
L 1
R 3
D 5
U 1
D 6
R 6
L 6
D 3
R 3
D 8
R 3
U 8
L 5
U 8
R 4
L 3
U 6
D 5
R 6
D 5
U 2
D 1
L 4
D 5
U 5
R 3
L 4
U 5
D 2
R 4
U 3
R 7
D 2
R 8
L 4
R 6
D 6
U 3
R 6
L 2
R 5
L 2
D 1
U 6
R 7
D 1
L 7
R 8
D 6
U 2
R 5
U 4
L 6
R 5
U 3
R 7
D 3
U 1
R 1
L 2
U 4
L 4
R 5
D 2
U 8
D 4
R 6
L 1
R 3
U 3
L 7
D 8
U 8
D 3
U 2
R 1
D 6
R 5
L 8
U 7
L 8
D 7
U 5
L 6
U 3
R 3
U 8
D 1
U 3
D 6
R 7
U 6
R 4
U 7
L 8
D 4
L 5
R 3
L 1
D 1
R 6
L 4
D 8
R 1
L 6
D 8
L 2
U 6
D 9
L 4
R 6
D 8
U 4
D 5
L 2
R 9
L 3
D 1
R 4
L 8
R 7
U 9
R 1
D 7
R 5
L 4
R 3
L 5
U 4
R 9
U 7
D 8
L 5
U 7
R 2
D 1
U 6
R 1
L 5
D 3
U 3
L 1
D 5
R 3
L 8
D 7
U 3
D 7
R 2
U 3
R 1
D 1
R 3
L 1
D 3
L 3
R 3
D 3
L 9
U 6
L 6
R 6
U 2
R 1
L 6
D 4
R 2
U 9
R 1
L 3
R 7
U 6
R 8
U 6
R 1
U 5
D 5
R 6
L 6
U 5
L 1
D 4
U 3
D 5
R 1
U 9
D 1
R 1
L 7
R 7
L 5
R 1
U 6
R 4
D 8
R 2
L 2
D 1
U 1
R 7
D 6
L 1
R 8
U 2
R 4
U 8
R 1
D 2
U 3
R 5
U 8
R 1
L 6
U 4
L 4
R 1
D 1
L 3
R 8
L 2
D 9
U 4
L 3
U 9
L 9
D 6
U 8
R 3
L 2
U 4
L 1
D 8
L 8
R 5
L 1
U 3
L 6
D 5
U 9
D 4
U 2
D 2
U 1
R 6
D 1
R 1
D 8
L 3
D 10
R 5
D 6
L 1
U 4
D 6
R 6
U 4
L 10
D 1
L 3
R 9
U 5
L 5
R 1
U 7
R 6
L 10
R 1
U 10
D 3
R 1
L 1
R 7
U 4
R 3
U 7
R 2
D 4
R 2
L 2
R 9
L 8
U 10
L 9
D 4
L 2
D 9
L 10
D 2
U 5
L 7
R 1
D 1
L 5
D 8
L 3
R 4
D 5
R 4
U 4
L 8
D 9
L 10
R 5
L 5
R 8
D 5
L 3
R 2
L 2
R 1
L 5
U 7
D 7
L 10
D 1
L 9
D 7
R 1
U 7
R 10
L 7
U 8
D 8
L 9
R 9
D 3
L 1
D 6
U 11
D 5
U 5
L 6
R 1
U 11
D 3
U 7
L 7
R 9
L 1
R 6
L 2
U 6
D 7
L 8
U 10
R 2
D 9
L 11
R 7
L 3
U 7
D 3
R 1
D 6
L 5
D 4
R 7
L 8
U 11
L 10
D 2
L 3
R 4
D 1
R 6
L 1
R 7
L 3
R 4
U 5
L 4
U 8
D 2
R 3
U 7
R 8
L 2
R 5
U 9
R 1
U 2
R 8
L 4
R 1
L 10
U 6
L 3
R 7
U 1
L 8
U 2
D 1
L 5
R 11
U 5
L 7
U 6
L 7
U 4
D 10
R 11
L 2
R 8
D 1
R 8
D 1
U 6
L 7
R 6
U 3
R 8
D 10
L 8
D 3
R 1
U 5
L 3
U 7
R 7
D 10
L 7
U 11
L 9
D 2
L 6
R 6
L 6
U 11
D 4
L 10
R 11
D 11
L 1
D 1
U 7
R 10
U 3
L 5
R 9
U 1
L 12
U 9
R 8
L 12
U 9
L 11
R 11
D 11
R 12
L 9
D 2
U 9
L 1
R 5
D 2
U 10
R 2
U 9
D 6
R 4
L 3
U 4
L 4
D 8
U 9
R 3
D 7
L 11
U 11
D 5
R 4
L 5
D 11
U 8
L 5
D 2
L 5
D 1
R 9
L 1
U 1
D 11
R 2
D 10
L 6
D 8
U 12
D 9
L 3
R 2
U 3
D 2
L 6
D 7
U 1
L 4
U 10
D 12
R 9
D 7
U 6
L 8
D 5
R 4
U 6
L 4
R 5
D 4
U 3
L 10
D 7
U 1
L 2
R 8
U 10
L 11
R 9
D 5
U 10
L 10
R 10
L 5
R 3
D 12
U 11
R 1
U 2
L 6
D 9
U 5
L 10
D 9
R 9
D 10
L 12
R 3
D 3
L 7
D 4
U 5
L 2
D 1
R 11
U 8
L 7
D 10
L 5
U 8
R 4
D 12
L 10
U 1
R 11
D 4
L 12
D 7
R 1
D 12
U 3
R 2
D 3
U 9
D 9
R 2
D 3
L 12
D 10
U 3
R 8
L 9
R 6
L 11
R 3
L 7
R 10
D 3
R 2
D 3
U 8
L 7
U 2
D 4
R 6
U 8
R 4
L 5
U 6
D 1
U 1
R 4
D 5
U 3
D 13
U 4
D 8
U 3
L 12
U 7
D 7
L 2
R 13
L 11
U 8
L 7
R 11
U 13
L 5
R 10
D 2
R 5
D 11
R 8
U 9
D 13
R 4
L 4
D 1
R 6
L 4
U 4
D 5
L 13
D 10
R 10
D 1
U 8
D 8
L 9
U 12
R 9
L 10
R 1
U 12
L 3
D 10
L 9
R 12
D 3
U 1
D 13
L 6
D 13
L 13
D 12
L 7
R 2
U 7
D 7
R 7
U 12
R 6
D 3
R 12
L 13
D 12
U 6
D 1
R 2
U 6
L 12
R 7
D 2
U 13
D 2
U 8
R 10
D 13
L 1
R 10
D 6
U 8
R 13
U 4
R 2
L 1
D 5
R 6
U 13
R 3
L 4
D 14
L 8
D 11
U 13
D 7
L 1
U 2
L 13
R 9
L 10
R 8
L 4
U 12
D 9
U 8
D 3
U 14
D 1
L 7
R 1
D 1
R 3
L 10
R 6
L 6
D 14
U 9
D 9
R 11
D 14
L 4
R 9
L 4
D 1
U 7
D 14
U 14
L 2
D 1
L 14
D 11
R 9
L 11
D 2
R 8
L 2
R 2
U 13
D 1
L 4
U 6
D 6
R 1
D 10
R 9
L 9
R 1
D 13
U 4
L 7
D 7
R 2
L 2
U 11
R 8
U 10
D 9
R 6
U 8
D 9
U 3
R 2
D 8
U 4
R 7
L 10
R 3
U 7
D 10
R 4
U 6
L 9
D 13
L 6
U 11
D 12
U 10
L 14
U 9
D 3
L 6
R 5
U 2
R 15
D 1
R 8
D 3
U 10
D 5
U 3
R 6
D 8
U 15
L 13
D 6
L 13
U 5
D 5
L 6
D 7
U 12
L 14
D 14
U 14
R 6
U 9
R 9
L 15
U 6
R 8
U 6
D 6
L 3
D 10
L 9
D 11
L 12
D 9
U 12
D 6
L 10
R 6
L 10
U 9
D 4
R 3
D 13
R 6
D 6
U 8
D 2
L 2
R 12
D 4
L 14
U 4
L 9
U 10
R 13
D 13
L 9
D 3
U 15
D 14
L 7
U 2
D 5
L 14
U 14
D 4
L 12
U 4
R 3
U 7
D 5
U 8
D 6
R 6
D 13
L 6
R 6
U 8
L 6
U 6
D 7
U 7
R 5
U 15
D 11
L 12
R 4
L 9
U 14
R 6
L 2
U 5
L 5
U 6
L 5
U 9
R 11
D 11
L 2
U 12
L 6
D 6
U 5
L 7
U 3
L 10
D 12
R 16
L 8
U 7
D 11
U 5
R 5
D 10
R 11
D 15
U 8
D 14
U 16
R 7
D 9
U 9
D 14
U 4
D 11
U 5
R 7
D 11
R 4
U 4
R 11
D 3
U 9
D 11
R 13
U 9
D 2
L 6
D 12
R 16
L 7
U 13
R 14
L 7
U 11
R 14
L 11
D 14
R 2
L 14
U 12
D 4
L 2
D 8
L 5
D 7
U 3
R 10
D 12
R 4
L 16
U 1
R 9
U 14
L 15
D 3
L 5
U 10
L 6
U 9
D 10
R 3
L 13
R 8
D 6
R 9
U 14
L 14
D 8
R 7
D 4
U 1
L 4
U 13
R 12
L 2
D 1
U 2
R 5
L 13
D 12
R 10
L 1
D 3
L 4
R 1
D 1
R 12
L 3
U 4
L 12
R 14
L 7
R 7
D 12
U 1
D 15
U 6
R 2
U 6
R 8
L 3
U 8
D 5
R 11
U 12
R 14
U 8
R 14
U 4
R 15
D 17
U 8
L 7
U 2
L 10
D 11
U 12
R 10
L 11
U 2
L 10
D 11
L 3
R 12
D 5
U 3
L 12
D 12
L 2
U 2
R 14
U 2
D 5
R 6
L 14
D 8
R 10
D 14
L 14
U 5
D 8
L 1
U 3
D 8
L 4
U 4
R 16
D 2
L 14
R 10
L 13
D 8
U 12
L 4
R 14
L 16
D 12
L 15
R 8
U 2
D 8
L 14
U 5
D 7
U 17
L 10
R 8
D 11
R 6
D 6
U 10
L 16
D 3
U 10
L 1
R 14
D 6
R 10
D 1
R 4
U 16
L 17
U 3
D 11
U 12
D 8
U 16
L 7
U 10
L 2
U 14
D 15
L 4
U 1
R 13
L 12
D 15
U 6
R 17
L 15
R 15
D 16
U 12
D 4
L 4
R 7
D 13
L 11
U 7
L 4
R 10
D 15
L 15
U 14
L 5
U 2
L 10
R 17
U 7
D 11
L 11
D 4
R 15
D 6
U 7
R 7
U 10
L 5
U 15
L 2
U 1
D 6
U 11
L 13
D 17
R 13
U 6
R 12
D 4
L 15
R 11
U 14
R 17
D 18
L 16
D 9
U 9
R 8
L 17
U 9
R 8
D 5
U 7
L 5
R 18
L 17
U 13
L 14
R 2
D 15
U 1
R 10
D 13
R 14
D 7
U 13
L 18
U 6
R 13
U 16
L 2
R 1
L 17
R 9
U 3
R 7
D 10
L 3
D 15
L 13
R 4
D 16
U 7
D 6
R 8
U 3
L 4
U 14
D 3
L 17
R 3
L 18
D 7
L 1
U 8
D 1
U 9
L 17
U 6
L 5
R 10
L 18
U 2
D 18
U 1
D 13
U 11
R 9
U 17
D 14
U 5
R 3
U 11
D 1
L 2
R 5
U 17
R 12
U 5
D 4
U 14
D 6
U 7
L 18
D 4
R 17
D 6
U 8
R 8
L 5
R 6
L 4
R 1
U 9
R 17
U 11
L 5
R 17
L 11
D 9
L 11
D 16
R 11
D 17
L 13
D 11
L 18
R 13
L 1
D 19
L 1
U 15
D 19
U 13
L 18
U 19
R 8
U 9
L 17
R 6
L 12
R 19
D 15
R 15
U 10
D 3
L 10
D 18
U 18
R 6
U 13
R 6
D 17
R 12
D 16
R 3
D 17
U 12
D 9
R 6
U 12
L 10
U 8
D 4
U 13
R 1
U 4
L 11
D 1
U 13
D 13
L 12
D 6
U 18
L 11
D 16
R 5
D 2
U 19
R 12
U 18
L 5
R 12
L 3
U 9
R 13
L 1
U 10
R 16
L 6
U 11
L 3
R 6
U 5
R 18
U 8
R 13
D 19
R 14
U 14
R 3
D 18
L 18
U 5
L 8
U 13
R 11
U 12
D 19
L 10
R 13
D 18
R 16
L 14
D 13
U 2
D 7
U 3
L 8
R 15
D 19`
}