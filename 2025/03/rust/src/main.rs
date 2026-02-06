use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

fn part1(num_rows: &Vec<Vec<u64>>) {
    let mut biggest_nums: Vec<u64> = Vec::new();

    for row in num_rows {
        let mut max = 0;
        let mut a = 0;
        while a < (row.len() - 1) {
            let mut b = a + 1;
            while b < row.len() {
                let current: u64 = format!("{}{}", row[a], row[b]).parse().expect("virhe pppp");
                // println!("{:?}, {:?} = {:?}, {:?}", a, b, row[a], row[b]);
                if current > max {
                    max = current;
                }
                b += 1;
            }
            a += 1;
        }
        biggest_nums.push(max);
    }
    
    let sum: u64 = biggest_nums.iter().sum();
    // println!("{:?}", biggest_nums);
    println!("{:?}", sum);
}
fn part2(num_rows: &Vec<Vec<u64>>) {
/*
    987654321111111
    811111111111119
    234234234234278
    818181911112111
    
    5 - 3
    23423
*/
    let mut biggest_nums: Vec<u64> = Vec::new();
    let num_size = 12;

    for row in num_rows {
        let mut found: Vec<u64> = Vec::new();
        let mut start_index = 0;
        // find biggest chars in row
        while found.len() < num_size {
            let mut biggest: u64 = 0;
            let mut i = start_index;
            // find biggest char
            while i <= row.len() - (num_size - found.len()) {
                if row[i] > biggest {
                    biggest = row[i];
                    start_index = i + 1;
                }
                i += 1;
            }
            found.push(biggest);
        }
        // println!("found: {:?}", found);
        let mut num_str = String::new();
        for e in found {
            let p = &e.to_string();
            num_str.push_str(p);
        }
        // println!("num_str: {:?}", num_str);
        
        let num: u64 = num_str.parse().expect("virhe löydetyn numberon parsinnassa");
        biggest_nums.push(num);
    }

    
    let sum: u64 = biggest_nums.iter().sum();
    // println!("{:?}", biggest_nums);
    println!("sum: {:?}", sum);
}

fn main() {
    // let text = read_input("../example.txt");
    let text = read_input("../input.txt");
   
    let num_rows: Vec<Vec<u64>> = text.split('\n')
        .filter(|x| x.len() > 0)
        .map(|row| row
            .split("")
            .filter(|x| x.len() > 0)
            .map(|x| x.parse().expect("virhe parsinnassa"))
            .collect()
        ).collect();
    // println!("{:?}", cc);
   
    
    println!("part1 ========================================");
    part1(&num_rows);
    println!("part2 ========================================");
    part2(&num_rows);

}
