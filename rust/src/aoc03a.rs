use std::str::FromStr;

use regex::Regex;

use crate::util;

fn aoc03a() -> i32 {
    let contents = util::read_file("../inputs/03a.txt");
    let re = Regex::new(r"mul\(\d{1,3},\d{1,3}\)").unwrap();

    let mut total = 0;
    for c in re.captures_iter(&contents) {
        let (s, _): (&str, [&str; 0]) = c.extract();
        total += multiply(s);
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
    fn test_aoc03a() {
        assert_eq!(173517243, aoc03a())
    }
}
