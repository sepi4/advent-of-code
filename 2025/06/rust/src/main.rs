use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

fn part1(rows: &Vec<&str>) -> isize {
    let ss: Vec<Vec<&str>> = rows.iter().map(|row| row.split_whitespace().collect()).collect();
    
    let mut result: isize = 0;
    for x in 0..ss[0].len() {
        let mut inner_arr: Vec<isize> = Vec::new();
        for y in 0..ss.len() - 1 {
            let v = ss[y][x];
            // println!("{:?}", v);
            let num = v.parse().expect("virhe parsinnassa");
            inner_arr.push(num);
        }
        let mut inner_result;
        // result += inner_arr.iter().sum::<isize>();
        let op = ss[ss.len() - 1][x];
        if op == "*" {
            // mult
            inner_result = 1;
            for v in inner_arr {
                inner_result *= v;
            }
        } else {
            // sum
            inner_result = 0;
            for v in inner_arr {
                inner_result += v;
            }
        }
        result += inner_result;
    }
    println!("{:?}", result);
    return result;
}

fn part2(rr: &Vec<&str>) -> isize {
    // tehdään char 2d matriisi 
    let rows: Vec<Vec<char>> = rr.iter().map(|x| x.chars().collect()).collect();

    // loopataan läpi ylhäältä -> alas ja vasemmalta -> oikealle
    // 
    // kun törmää tyhjään merkiin
    //      - jos ei ollut numeroa muistissa, ei tehdä mitään
    //      - jos oli numero muistissa, laitetaan numero listaan1 (inttinä)
    // kun törmää numeroon 
    //      - jos ei ollut numeroa muistissa, laitetaan muistiin (stringinä)
    //      - jos oli numero muistissa, laitetaan sen perään toinen numero stringinä (numero pitenee)
    // kun törmää operaatio merkkiin, 
    //      - tehdään operaatio kaikilla numerolla listassa ja tallennetaan se listaan2. 
    //      - tyhjennetään lista1
    
    let mut x = rows[0].len() - 1;
    let mut result_nums: Vec<isize> = Vec::new();
    let mut inner_nums: Vec<isize> = Vec::new();
    loop {
        let mut num_str = String::new();
        
        for y in 0..rows.len() {
            let v: char = rows[y][x];
            if v == ' ' {
                if num_str.len() > 0 {
                    inner_nums.push(num_str.parse().expect("virhe parsinnassa "));
                    num_str.clear();
                }
            } else if v == '*' || v == '+' {
                // operaatiot
                if num_str.len() > 0 {
                    inner_nums.push(num_str.parse().expect("virhe parsinnassa "));
                    num_str.clear();
                }
                let mut value = 0;
                match v {
                    '*' => {
                        value = 1;
                        for n in &inner_nums {
                            value *= n;
                        }
                    },
                    '+' => {
                        for n in &inner_nums {
                            value += n;
                        }
                    },
                    _ => panic!("not '*' nor '+'"),
                }
                result_nums.push(value);
                inner_nums.clear();
            } else {
                // numero
                num_str.push(v);
            }
        }
        if x == 0 {
            break;
        }
        x -= 1;
    }
    let result = result_nums.iter().sum::<isize>();
    println!("{:?}", result);
    return result;
}

fn main() {
    // let text = read_input("../example.txt");
    let text = read_input("../input.txt");
    
    let rows: Vec<&str> = text.split("\n").filter(|s| s.len() > 0).collect();
    // println!("{:?}", rows);
    

    
    println!("part1 ================================");
    part1(&rows);
    println!("part2 ================================");
    part2(&rows);
}
