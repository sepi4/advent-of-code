const DIR: [(isize, isize); 8] = [
    (-1, -1),
    (-1, 0),
    (-1, 1),

    (0, -1),
    (0, 1),

    (1, -1),
    (1, 0),
    (1, 1),
];

fn read_input(file_path: &str) -> String {
    let text = std::fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

fn get_next_counter(rows: &Vec<Vec<char>>, ys: isize, xs: isize) -> isize {
    let mut counter = 0;

    // loop directions constant
    for (yd, xd) in DIR {
        let y = ys + yd;
        let x = xs + xd;
        // skip when some of indexes is going over the limits
        if y < 0 
        || x < 0 
        || y >= rows.len() as isize // can't compare isize and usize
        || x >= rows[0].len() as isize
        {
            continue;
        }
        let y = y as usize; // can index only with usize
        let x = x as usize;
        if rows[y][x] == '@' {
            counter += 1;
        }
    }
    return counter;
}

fn get_rows(text: String) -> Vec<Vec<char>> {
    let rows: Vec<Vec<char>> = text.lines()
        .map(|line| line.chars().collect())
        .collect();
    rows
}

fn part1(rows: &Vec<Vec<char>>) -> isize {
    let mut amount = 0;
    
    for y in 0..rows.len() {
        for x in 0..rows[0].len() {
            if rows[y][x] == '@' {
                let count = get_next_counter(rows, y as isize, x as isize);
                if count < 4 {
                    amount += 1;
                }
            }
        }
    }
    println!("amount: {:?}", amount);
    return amount;
}

fn remove_rolls(rows: &mut Vec<Vec<char>>) -> isize {
    let mut amount = 0;
    
    let mut to_be_removed: Vec<(usize, usize)> = vec![];
    for y in 0..rows.len() {
        for x in 0..rows[0].len() {
            if rows[y][x] == '@' {
                let count = get_next_counter(rows, y as isize, x as isize);
                if count < 4 {
                    amount += 1;
                    to_be_removed.push((y,x));
                }
            }
        }
    }
    for (y, x) in to_be_removed {
        rows[y][x] = '.';
    }
    amount
}

fn part2(rows: &mut Vec<Vec<char>>) {
    let mut amount = 0;
    loop {
        let n = remove_rolls(rows);
        if n == 0 {
            break;
        }
        amount += n;
    }
    println!("amount: {}", amount);
}

fn main() {
    // let text = read_input("../example.txt");
    let text = read_input("../input.txt");
    
    let mut rows: Vec<Vec<char>> = get_rows(text);

    println!("part1 =================");
    part1(&rows); // pass just reference
    
    println!("part2 =================");
    part2(&mut rows); // pass mutable reference
}
