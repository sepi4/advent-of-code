use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}


#[derive(Debug)]
struct Stream {
    coor: (usize, usize), // (y, x)
}

fn part2() {
}

fn remove_same_position(ss: &mut Vec<Stream>) {
    // println!("removing {:?}",ss);
    let mut i = 0;
    while i < ss.len() - 1 {
        let mut j = i + 1;
        while j < ss.len() {
            let (ay, ax) = ss[i].coor;
            let (by, bx) = ss[j].coor;
            if ay == by && ax == bx {
                ss.remove(j);
            }
            j += 1;
        }
        i += 1;
    }
    // println!("after removing {:?}",ss);
}

fn part1(cells: &Vec<Vec<char>>) {
    let mut count_splits = 0;
    // println!("{:?}", );
    let si = cells[0].iter().position(|c| *c == 'S');
    if si == None {
        panic!("not found starting 'S'");
    }
    let mut streams: Vec<Stream> = vec![Stream{coor: (0, si.unwrap())}];
    loop {
        let mut next_row: Vec<Stream> = vec![];
        for s in &mut streams {

            s.coor.0 += 1;

            if cells[s.coor.0][s.coor.1] == '^' {
                next_row.push(Stream{coor: (s.coor.0, s.coor.1 + 1)});
                s.coor.1 -= 1;
                count_splits += 1;
            }
        }
        if next_row.len() > 0 {
            streams.extend(next_row);
            remove_same_position(&mut streams);
        }
        
        // check if reached bottom
        if streams[0].coor.0 == cells.len() - 1 {
            break;
        }
    }
    println!("{:?}", streams.len());
    println!("splits {:?}", count_splits);
}


fn main() {
    let text = read_input("../example.txt");
    // let text = read_input("../input.txt");

    // println!("{:?}", text);
    let cells = text.split('\n')
        .filter(|row| !row.is_empty())
        .map(|row| row.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    
    // // println!("{:?}", cells);
    // for row in cells {
    //     println!("{:?}", row);
    // }

    println!("part1 ================================");
    part1(&cells);
    
    // println!("part2 ================================");
    // part2(&rows);
    

}
