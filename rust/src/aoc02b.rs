use std::{
    fs::File,
    io::{BufRead, BufReader},
    str::FromStr,
};

fn aoc02b() -> i32 {
    let f = File::open("../inputs/02.txt").expect("could not open file");
    let mut reader = BufReader::new(f);
    let mut buffer = String::new();

    let mut total = 0;
    loop {
        match reader.read_line(&mut buffer) {
            Ok(0) => {
                break;
            }
            Ok(_) => {
                if is_safe(&buffer.trim()) {
                    total += 1;
                }
                buffer.clear();
            }
            Err(e) => {
                eprintln!("Error reading line: {}", e);
                break;
            }
        };
    }

    total
}

fn is_safe(line: &str) -> bool {
    let mut items: Vec<i32> = vec![];
    line.split(" ")
        .for_each(|i| items.push(FromStr::from_str(i).expect("could parse into i32")));

    for i in (-1..items.len() as i32).step_by(1) {
        let mut l = items.clone();
        if i >= 0 {
            l.remove(i as usize);
        }

        if removed_item_check(l) {
            return true;
        }
    }

    false
}

fn removed_item_check(items: Vec<i32>) -> bool {
    let mut all_inc = true;
    let mut all_dec = true;
    for i in (1..items.len()).step_by(1) {
        let curr = items[i];
        let prev = items[i - 1];

        all_inc = all_inc && curr > prev;
        all_dec = all_dec && curr < prev;

        let diff = i32::abs(prev - curr);
        if diff < 1 || diff > 3 {
            return false;
        }
    }

    all_inc || all_dec
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_aoc02b() {
        assert_eq!(621, aoc02b());
    }
}
