use std::{str::FromStr, usize};

use regex::Regex;

use crate::util;

fn aoc03b() -> i32 {
    let contents = util::read_file("../inputs/03a.txt");
    let re = Regex::new(r"mul\(\d{1,3},\d{1,3}\)").unwrap();

    let mut total = 0;
    let mut full_start = 0;
    let mut enabled = true;
    for c in re.find_iter(&contents) {
        let start = c.start();
        let end = c.end();
        let before = &contents[full_start..start];

        let d = before.rfind("do()").map(|i| i as i32).unwrap_or(-1);
        let dd = before.rfind("don't()").map(|i| i as i32).unwrap_or(-1);

        if dd > d {
            enabled = false;
            full_start = dd as usize + "don't()".len();
        } else if d > dd {
            enabled = true;
            full_start = d as usize + "do()".len();
        }

        if enabled {
            total += multiply(&contents[start..end]);
        }
    }

    total
}

fn multiply(s: &str) -> i32 {
    let all: Vec<&str> = s.split(",").collect();
    let left: i32 = FromStr::from_str(all[0].strip_prefix("mul(").unwrap()).unwrap();
    let right: i32 = FromStr::from_str(all[1].strip_suffix(")").unwrap()).unwrap();

    left * right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_aoc03b() {
        assert_eq!(100450138, aoc03b())
    }
}
